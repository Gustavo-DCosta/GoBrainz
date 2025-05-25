# GoBrainz – The Nerdy Gopher Assistant for Notion

<p align="center">
  <img src="./README/Logo de GoBeanz et gopher.png" alt="GoBeanz mascot"/>
</p>

**GoBrainz** is a minimal CLI productivity assistant powered by the Notion API and written in Go.
No web fluff. Just brains, beans, and business.

> 💡 *“Let your projects flow like coffee-fueled commands.”*

---

## ⚙️ Features

> Talk to Notion like a dev talks to a terminal.

*  **Pages**: Create, update, and retrieve page content.
*  **Databases**: Manage properties, entries, and schemas.
*  **Users**: Access user profiles & permissions.
*  **Comments**: Read & write page comments.
*  **Search**: Query across workspace content.
*  **OAuth 2.0**: Secure authentication flow.
*  **Link Previews**: Customize how links look when shared.

---

## 🚀 Getting Started

```bash
go install github.com/Gustavo-DCosta/GoBrainz@latest
```

Then just run it:

```bash
gobrainz
```

---

## 🎨 Choose Your Syntax Style

Upon launch, GoBrainz lets you choose your preferred CLI syntax for maximum ✨vibes✨:

```go
color.Cyan("Welcome to GoBrainz, your nerdy Notion sidekick.")

color.Magenta("Before starting, would you like to select the syntax you are most comfortable with?")
color.Yellow(`
1. @argument
2. argument?
3. GoBrainz argument
`)

fmt.Scan(&choice)

for {
  switch(choice) {
    case 1:
      file, _ := os.Create("syntax.txt")
      defer file.Close()
      break
    case 2:
      file, _ := os.Create("syntax.txt")
      defer file.Close()
      break
    case 3:
      file, _ := os.Create("syntax.txt")
      defer file.Close()
      break
  }
}
```

---

## Future Plans

* 🤖 AI + Notion: Add GPT-assisted database and project creation.
* 🧠 Auto-generated templates based on task type.
* 🔄 Live sync w/ Notion and local files (markdown, .env, etc).
* 🌐 Optional Web UI? (maybe?? if GoFiber doesn’t murder me)

---

## Auth Note

To use the Notion API, make sure to set your integration secret and follow OAuth setup steps [here](https://developers.notion.com/docs/getting-started).

---

## Built With

* Go (Golang 🐹)
* Notion API
* Cobra CLI (planned)
* Color library for cool terminal vibes

---

## License

MIT – steal it, fork it, improve it. Just give the gopher some love.

---

![Go Version](https://img.shields.io/badge/go-1.24-blue)
![Notion API](https://img.shields.io/badge/notion-api-orange)
![CLI-Lover](https://img.shields.io/badge/CLI-approved-brightgreen)
