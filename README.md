# vibecheck

Clearly tell other developers how vibecoded your project or PR is.

![vibe](https://img.shields.io/badge/vibe-10%25-white) ![AI](https://img.shields.io/badge/AI-90%25-red?logo=anthropic)

## What do the numbers mean?

**AI**: What percent of the code was ai generated. Did ai write 80% of the code? Then set `--ai 80`.

**vibes**: How vibecoded is this? This tells others what your confidence level in the code is. More vibe = less confident.

| vibe | Meaning |
|--|----|
| 100 | This whole thing is straight yolo, I haven't read it |
| 0 | I understand every line of this code as if I wrote it myself |
| 20 | I wrote or reviewed most of it, some pieces I just let the ai handle |

## Quickstart

1. Download from releases page

2. `chmod +x vibecheck`

3. Run

```
vibecheck --help
A CLI tool to generate GitHub PR body text based on vibe coding metrics

Usage:
  vibecheck [flags]

Flags:
      --agent string      AI agent used (amp, claude, cursor) (default "ai")
      --ai int            AI percentage (0-100)
      --artifact string   Artifact identifier (default "unknown")
  -h, --help              help for vibecheck
      --output string     Output format (text, json, github-markdown) (default "text")
      --vibes int         Vibe percentage (0-100) [REQUIRED]

```

## Examples

```
./vibecheck --agent claude --ai 90 --vibes 10 --output github-markdown
```
![vibe](https://img.shields.io/badge/vibe-10%25-white) ![AI](https://img.shields.io/badge/AI-90%25-red?logo=anthropic)

```
./vibecheck --agent amp --ai 25 --vibes 45 --output github-markdown --artifact T-007dd158-76eb-4daf-967f-8c6d683cde7b
```
![vibe](https://img.shields.io/badge/vibe-45%25-orange) ![AI](https://img.shields.io/badge/AI-25%25-yellow?logo=data:image/svg%2bxml;base64,PHN2ZyB3aWR0aD0iMjEiIGhlaWdodD0iMjEiIHZpZXdCb3g9IjAgMCAyMSAyMSIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTMuNzY4NzkgMTguMzAxNUw4LjQ5ODM5IDEzLjUwNUwxMC4yMTk2IDIwLjAzOTlMMTIuNzIgMTkuMzU2MUwxMC4yMjg4IDkuODY3NDlMMC44OTA4NzYgNy4zMzg0NEwwLjIyNTk0IDkuODkzMzFMNi42NTEzNCAxMS42Mzg4TDEuOTQxMzggMTYuNDI4MkwzLjc2ODc5IDE4LjMwMTVaIiBmaWxsPSIjRjM0RTNGIi8+CjxwYXRoIGQ9Ik0xNy40MDc0IDEyLjc0MTRMMTkuOTA3OCAxMi4wNTc1TDE3LjQxNjcgMi41Njg5N0w4LjA3ODczIDAuMDM5OTI0Nkw3LjQxMzggMi41OTQ4TDE1LjI5OTIgNC43MzY4NUwxNy40MDc0IDEyLjc0MTRaIiBmaWxsPSIjRjM0RTNGIi8+CjxwYXRoIGQ9Ik0xMy44MTg0IDE2LjM4ODNMMTYuMzE4OCAxNS43MDQ0TDEzLjgyNzYgNi4yMTU4OEw0LjQ4OTcxIDMuNjg2ODNMMy44MjQ3NyA2LjI0MTcxTDExLjcxMDEgOC4zODM3NkwxMy44MTg0IDE2LjM4ODNaIiBmaWxsPSIjRjM0RTNGIi8+Cjwvc3ZnPgo=) [![amp thread](https://img.shields.io/badge/amp_thread-T--007dd158--76eb--4daf--967f--8c6d683cde7b-green?logo=data:image/svg%2bxml;base64,PHN2ZyB3aWR0aD0iMjEiIGhlaWdodD0iMjEiIHZpZXdCb3g9IjAgMCAyMSAyMSIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTMuNzY4NzkgMTguMzAxNUw4LjQ5ODM5IDEzLjUwNUwxMC4yMTk2IDIwLjAzOTlMMTIuNzIgMTkuMzU2MUwxMC4yMjg4IDkuODY3NDlMMC44OTA4NzYgNy4zMzg0NEwwLjIyNTk0IDkuODkzMzFMNi42NTEzNCAxMS42Mzg4TDEuOTQxMzggMTYuNDI4MkwzLjc2ODc5IDE4LjMwMTVaIiBmaWxsPSIjRjM0RTNGIi8+CjxwYXRoIGQ9Ik0xNy40MDc0IDEyLjc0MTRMMTkuOTA3OCAxMi4wNTc1TDE3LjQxNjcgMi41Njg5N0w4LjA3ODczIDAuMDM5OTI0Nkw3LjQxMzggMi41OTQ4TDE1LjI5OTIgNC43MzY4NUwxNy40MDc0IDEyLjc0MTRaIiBmaWxsPSIjRjM0RTNGIi8+CjxwYXRoIGQ9Ik0xMy44MTg0IDE2LjM4ODNMMTYuMzE4OCAxNS43MDQ0TDEzLjgyNzYgNi4yMTU4OEw0LjQ4OTcxIDMuNjg2ODNMMy44MjQ3NyA2LjI0MTcxTDExLjcxMDEgOC4zODM3NkwxMy44MTg0IDE2LjM4ODNaIiBmaWxsPSIjRjM0RTNGIi8+Cjwvc3ZnPgo=)](https://ampcode.com/threads/T-007dd158-76eb-4daf-967f-8c6d683cde7b)

```
./vibecheck --ai 100 --vibes 100 --agent cursor --output github-markdown
```
![vibe](https://img.shields.io/badge/vibe-100%25-green) ![AI](https://img.shields.io/badge/AI-100%25-red?logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB2aWV3Qm94PSIwLDAsMjU2LDI1NiIgd2lkdGg9IjUwcHgiIGhlaWdodD0iNTBweCIgZmlsbC1ydWxlPSJub256ZXJvIj48ZyBmaWxsPSIjZmZmZmZmIiBmaWxsLXJ1bGU9Im5vbnplcm8iIHN0cm9rZT0ibm9uZSIgc3Ryb2tlLXdpZHRoPSIxIiBzdHJva2UtbGluZWNhcD0iYnV0dCIgc3Ryb2tlLWxpbmVqb2luPSJtaXRlciIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBzdHJva2UtZGFzaGFycmF5PSIiIHN0cm9rZS1kYXNob2Zmc2V0PSIwIiBmb250LWZhbWlseT0ibm9uZSIgZm9udC13ZWlnaHQ9Im5vbmUiIGZvbnQtc2l6ZT0ibm9uZSIgdGV4dC1hbmNob3I9Im5vbmUiIHN0eWxlPSJtaXgtYmxlbmQtbW9kZTogbm9ybWFsIj48ZyB0cmFuc2Zvcm09InNjYWxlKDUuMTIsNS4xMikiPjxwYXRoIGQ9Ik0yNC45NDkyMiwyLjk5ODA1Yy0wLjE2NiwwLjAwOTExIC0wLjMyNzEsMC4wNTk0NiAtMC40Njg3NSwwLjE0NjQ4bC0xOC4wMDU4NiwxMS4wMDE5NWMtMC4yOTczLDAuMTgxNzIgLTAuNDc4NTgsMC41MDUwOCAtMC40Nzg1MiwwLjg1MzUybDAuMDAzOTEsMjBjMC4wMDA0NiwwLjAzNzg4IDAuMDAzMDcsMC4wNzU3IDAuMDA3ODEsMC4xMTMyOGMwLjAwMzU5LDAuMDMyNzkgMC4wMDg4MSwwLjA2NTM4IDAuMDE1NjMsMC4wOTc2NmMwLjAwNjQ4LDAuMDI5NjQgMC4wMTQzLDAuMDU4OTYgMC4wMjM0NCwwLjA4Nzg5YzAuMDAwNjMsMC4wMDMyNiAwLjAwMTI5LDAuMDA2NTEgMC4wMDE5NSwwLjAwOTc3YzAuMDA4NjIsMC4wMjcxIDAuMDE4NCwwLjA1MzgxIDAuMDI5MywwLjA4MDA4YzAuMDk2NzEsMC4yMzY5IDAuMjgwNTQsMC40Mjc3MyAwLjUxMzY3LDAuNTMzMmwxNy45MTQwNiwxMC45MzE2NGMwLjMyMDE0LDAuMTk1NjMgMC43MjI4MywwLjE5NTYzIDEuMDQyOTcsMGwxNy45NzQ2MSwtMTEuMDA1ODZjMC4yOTczLC0wLjE4MTcyIDAuNDc4NTgsLTAuNTA1MDggMC40Nzg1MiwtMC44NTM1MmwtMC4wMDM5MSwtMTkuOTk0MTRjLTAuMDAwMTksLTAuMzQ3MzUgLTAuMTgwNjEsLTAuNjY5NzMgLTAuNDc2NTYsLTAuODUxNTZsLTE3Ljk5ODA1LC0xMS4wMDM5MWMtMC4xNzIyMSwtMC4xMDU3MiAtMC4zNzI0MSwtMC4xNTY3OSAtMC41NzQyMiwtMC4xNDY0OHpNMjUuOTk4MDUsNS43NzczNGwxNS4wNzgxMyw5LjIyMDdoLTE1LjA4Mzk4ek0yMy45OTgwNSw1Ljc4MTI1bC0wLjAwNTg2LDkuMjE0ODRoLTE1LjA3MDMxek0xMi4wMTU2MywxNi45OTYwOWgxMi43NjM2N2MwLjE0MDE0LDAuMDMxMjEgMC4yODUzNiwwLjAzMTg3IDAuNDI1NzgsMC4wMDE5NWgwLjAwMzkxaDE1LjA0NDkybC0xNC4yNTM5MSwyNC4zMTY0MWwtMC4wMDM5MSwtMTYuMzA4NTljLTAuMDAwMjEsLTAuMzY5ODcgLTAuMjA0NTUsLTAuNzA5NDIgLTAuNTMxMjUsLTAuODgyODF6TTcuOTk4MDUsMTcuMTI4OTFsMTUuMTIzMDUsOC4wMTU2M2wtMTUuMTIxMDksOC4xOTE0MXpNNDIsMTcuOTcwN3YxNC45MDAzOWwtNi42NTAzOSwtMy41NTI3M3pNMjMuOTk2MDksMjYuOTQ1MzFsMC4wMDU4NiwxNy4yNTc4MWwtMTQuOTg2MzMsLTkuMTQ0NTN6TTM0LjMzNTk0LDMxLjA0Njg4bDcuMDUwNzgsMy43NjM2N2wtMTQuNDM3NSw4LjgzOTg0eiI+PC9wYXRoPjwvZz48L2c+PC9zdmc+)
