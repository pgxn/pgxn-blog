package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const URL = "https://blog.pgxn.org/rss"

type Post struct {
	Title   string
	Summary string
	Slug    string
	Date    string
	ID      string `json:"id_string"`
	URL     string `json:"post_url"`
	Tags    []string
}

func (p Post) GetTitle() string {
	if p.Title == "" {
		p.Title = Truncate(p.Summary, 50)
		if len(p.Tags) == 1 && p.Tags[0] == "tweet" {
			p.Title += " ⓣ"
		}
	}
	if strings.Contains(p.Title, ":") {
		return `"` + p.Title + `"`
	}
	return p.Title
}

func (p Post) Year() string {
	year, _, _ := strings.Cut(p.Date, "-")
	return year
}

func (p Post) GetDate() string {
	return strings.Replace(strings.Replace(p.Date, " ", "T", 1), " GMT", "Z", 1)
}

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: main.go JSON_DIRECTORY MARKDOWN_DIRECTORY DEST_DIR")
	}

	jsonDir := os.Args[1]
	mdDir := os.Args[2]
	destDir := os.Args[3]
	created := 0
	skipped := 0
	tweets := 0

	entries, err := os.ReadDir(jsonDir)
	if err != nil {
		log.Fatal(err)
	}

	// tags := []string{}

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(path.Join(jsonDir, e.Name()))
		if err != nil {
			log.Fatal(err)
		}
		var post Post
		err = json.Unmarshal(data, &post)
		if err != nil {
			log.Fatal(err)
		}

		// Skip cross-posted tweets.
		if len(post.Tags) == 1 && post.Tags[0] == "tweet" {
			tweets++
		}

		mdPath := path.Join(mdDir, fmt.Sprintf("%s.html.md", post.ID))
		mdFile, err := os.Open(mdPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				skipped++
				continue
			}
			log.Fatal(err)
		}
		defer mdFile.Close()

		year := post.Year()
		yearDir := path.Join(destDir, year)
		err = os.MkdirAll(yearDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		destFile, err := os.Create(path.Join(yearDir, post.Slug+".md"))
		if err != nil {
			log.Fatal(err)
		}
		defer destFile.Close()

		// tags = append(tags, makeTags(post.Tags)...)

		fmt.Fprintf(destFile, "---\n")
		fmt.Fprintf(destFile, "title: %s\n", post.GetTitle())
		fmt.Fprintf(destFile, "slug: %s\n", post.Slug)
		fmt.Fprintf(destFile, "date: %s\n", post.GetDate())
		fmt.Fprintf(destFile, "aliases: [%s, /post/%s]\n", strings.TrimPrefix(post.URL, "https://blog.pgxn.org"), post.ID)
		fmt.Fprintf(destFile, "tags: [%s]\n", strings.Join(makeTags(post.Tags), ", "))
		fmt.Fprint(destFile, "---\n")

		scanner := bufio.NewScanner(mdFile)
		for scanner.Scan() {
			switch {
			case strings.HasPrefix(scanner.Text(), "# "):
				continue
			case strings.HasPrefix(scanner.Text(), "::: caption"):
				destFile.Write([]byte("<figure>\n<figcaption>\n"))
				// Skip to next :::
				for scanner.Scan() {
					if strings.HasPrefix(scanner.Text(), ":::") {
						destFile.Write([]byte("</figcaption>\n</figure>\n"))
						break
					}
					destFile.Write(scanner.Bytes())
					destFile.Write([]byte("\n"))
				}
			case strings.HasPrefix(scanner.Text(), ":::"):
				// Skip to next :::
				for scanner.Scan() {
					if strings.HasPrefix(scanner.Text(), ":::") {
						fmt.Fprint(destFile, "\n")
						break
					}
				}
			default:
				destFile.Write(scanner.Bytes())
				destFile.Write([]byte("\n"))
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		created++
	}

	// slices.Sort(tags)
	// tags = slices.Compact(tags)
	// for _, t := range tags {
	// 	fmt.Println(t)
	// }

	fmt.Printf("%v migrated (%v tweets); %v skipped\n", created, tweets, skipped)
}

func makeTags(cats []string) []string {
	caser := cases.Title(language.AmericanEnglish)
	tags := make([]string, len(cats))
	for i, c := range cats {
		switch c {
		case "fti", "FTI":
			tags[i] = "FTI"
		case "pgxn":
			tags[i] = "PGXN"
		case "pgxn client":
			tags[i] = "PGXN Client"
		case "pgxn utils":
			tags[i] = "PGXN Utils"
		case "html":
			tags[i] = "HTML"
		case "plperl":
			tags[i] = "PL/Perl"
		case "plpgsql":
			tags[i] = "PL/pgSQL"
		case "json":
			tags[i] = "JSON"
		case "jsonp":
			tags[i] = "JSONP"
		case "readme":
			tags[i] = "README"
		case "jfdi":
			tags[i] = "JFDI"
		case "jquery":
			tags[i] = "jQuery"
		case "http":
			tags[i] = "HTTP"
		case "httpd":
			tags[i] = "HTTPD"
		case "postgresql":
			tags[i] = "PostgreSQL"
		case "rfc":
			tags[i] = "RFC"
		case "mime":
			tags[i] = "MIME"
		case "mime type":
			tags[i] = "MIME Type"
		case "uri":
			tags[i] = "URI"
		case "uri templates":
			tags[i] = "URI Templates"
		case "user-submitted":
			tags[i] = "User Submitted"
		case "tls":
			tags[i] = "TLS"
		case "cpan":
			tags[i] = "CPAN"
		case "ssl":
			tags[i] = "SSL"
		case "javascript":
			tags[i] = "JavaScript"
		case "cpan meta spec":
			tags[i] = "CPAN Meta Spec"
		case "api":
			tags[i] = "API"
		case "api docs":
			tags[i] = "API Docs"
		case "oscon":
			tags[i] = "OSCON"
		case "api server":
			tags[i] = "API Server"
		case "ascii":
			tags[i] = "ASCII"
		case "pgday":
			tags[i] = "PgDay"
		case "pgwest":
			tags[i] = "PgWest"
		case "pgtap":
			tags[i] = "pgTAP"
		case "case-insensitive":
			tags[i] = "Case Insensitive"
		case "case-sensitive":
			tags[i] = "Case Sensitive"
		case "howto":
			tags[i] = "HOWTO"
		case "semver":
			tags[i] = "SemVer"
		case "pgxs":
			tags[i] = "PGXS"
		case "myyearbook", "MyYearbook":
			tags[i] = "myYearbook"
		case "GitHub Actions", "GitHub Workflows", "USE_PGXS", "pg_config",
			"myYearbook", "META.json", "PGXN Manager", "PGXN::Manager", "PgWest",
			"PostgreSQL", "PostgreSQL 9", "RDG", "REST", "TigerLead", "UE", "UI",
			"contrib", "depesz", "make", "metacpan", "mod_proxy", "mod_ssl",
			"pg_regress", "rsync", "sudo":
			tags[i] = c
		case "HTTP status codes":
			tags[i] = "HTTP Status Codes"
		case "listen-notify":
			tags[i] = "LISTEN/NOTIFY"
		case "travisci":
			tags[i] = "Travis CI"
		case "adaptive_estimator":
			tags[i] = "Adaptive Estimator"
		case "github":
			tags[i] = "GitHub"
		case "develoment":
			tags[i] = "Development"
		case "back-to-work":
			tags[i] = "Back to Work"
		default:
			tags[i] = caser.String(c)
		}
	}
	return tags
}

// Truncate truncates the string in s to the specified length.
func Truncate(text string, length int) string {
	text = strings.TrimSpace(text)
	if before, _, found := strings.Cut(text, "\n"); found {
		text = before
	}
	if before, _, found := strings.Cut(text, "!"); found {
		text = before
	}
	if before, _, found := strings.Cut(text, "http:"); found {
		text = before
	}
	if before, _, found := strings.Cut(text, "#"); found {
		text = before
	}
	if before, _, found := strings.Cut(text, ". "); found {
		text = before
	}
	if utf8.RuneCountInString(text) <= length {
		return strings.TrimRight(text, ".!:;? ")
	}

	var lastWordIndex, lastNonSpace, currentLen, endTextPos, nextTag int

	for i, r := range text {
		if i < nextTag {
			continue
		}

		currentLen++
		if unicode.IsSpace(r) {
			lastWordIndex = lastNonSpace
		} else if unicode.In(r, unicode.Han, unicode.Hangul, unicode.Hiragana, unicode.Katakana) {
			lastWordIndex = lastNonSpace
			lastNonSpace = i + utf8.RuneLen(r)
			if currentLen <= length {
				lastWordIndex = lastNonSpace
			}
		} else {
			lastNonSpace = i + utf8.RuneLen(r)
		}

		if currentLen > length {
			if lastWordIndex == 0 {
				endTextPos = i
			} else {
				endTextPos = lastWordIndex
			}
			var out strings.Builder
			out.WriteString(text[0:endTextPos])
			return strings.TrimRight(out.String(), ".!:;? ")
		}
	}

	return strings.TrimRight(text, ".!:;? ")
}
