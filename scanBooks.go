package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type NCUData struct {
	USBooks []Book `json:"us_books"`
	UKBooks []Book `json:"uk_books"`
}

type Book struct {
	BookName string `json:"bookName"`
	Units    []Unit `json:"units"`
}

type Unit struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	LRC   string `json:"lrc"`
	Audio string `json:"audio"`
}

func extractTitleFromFilename(filename string) string {
	base := strings.TrimSuffix(filename, ".lrc")
	base = strings.TrimSuffix(base, ".mp3")
	re := regexp.MustCompile(`^\d+(&\d+)?\.`)
	title := re.ReplaceAllString(base, "")
	title = strings.TrimSpace(title)
	return title
}

func scanUnits(bookDirPath, lang, bookName string) ([]Unit, error) {
	entries, err := os.ReadDir(bookDirPath)
	if err != nil {
		return nil, err
	}
	var units []Unit
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".lrc") {
			continue
		}
		slug := strings.TrimSuffix(name, ".lrc")
		title := extractTitleFromFilename(name)
		if title == "" {
			fmt.Printf("警告: 无法从文件名提取标题: %s\n", name)
			continue
		}
		relPath := filepath.Join(lang, bookName, name)
		relAudio := strings.TrimSuffix(relPath, ".lrc") + ".mp3"
		units = append(units, Unit{
			Title: title,
			Slug:  slug,
			LRC:   relPath,
			Audio: relAudio,
		})
	}
	sort.Slice(units, func(i, j int) bool {
		re := regexp.MustCompile(`^\d+`)
		numI := re.FindString(units[i].Slug)
		numJ := re.FindString(units[j].Slug)
		return numI < numJ
	})
	return units, nil
}

func scanLanguage(langDirPath, lang string) ([]Book, error) {
	entries, err := os.ReadDir(langDirPath)
	if err != nil {
		return nil, err
	}
	var books []Book
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		bookName := entry.Name()
		bookFullPath := filepath.Join(langDirPath, bookName)
		units, err := scanUnits(bookFullPath, lang, bookName)
		if err != nil {
			fmt.Printf("警告: 扫描 %s 失败: %v\n", bookFullPath, err)
			continue
		}
		if len(units) > 0 {
			books = append(books, Book{
				BookName: bookName,
				Units:    units,
			})
		}
	}
	sort.Slice(books, func(i, j int) bool {
		return books[i].BookName < books[j].BookName
	})
	return books, nil
}

func main() {
	usDir := "./us"
	ukDir := "./uk"

	data := NCUData{}

	if _, err := os.Stat(usDir); err == nil {
		books, err := scanLanguage(usDir, "us")
		if err != nil {
			fmt.Printf("扫描美音目录失败: %v\n", err)
		} else {
			data.USBooks = books
			fmt.Printf("✅ 美音: 发现 %d 本书籍\n", len(books))
		}
	} else {
		fmt.Println("美音目录不存在，跳过")
	}

	if _, err := os.Stat(ukDir); err == nil {
		books, err := scanLanguage(ukDir, "uk")
		if err != nil {
			fmt.Printf("扫描英音目录失败: %v\n", err)
		} else {
			data.UKBooks = books
			fmt.Printf("✅ 英音: 发现 %d 本书籍\n", len(books))
		}
	} else {
		fmt.Println("英音目录不存在，跳过")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("生成 JSON 失败: %v\n", err)
		return
	}
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}
	fmt.Println("🎉 成功生成 data.json")
}
