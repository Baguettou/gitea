	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/modules/analyze"
	file     *DiffFile
	language := ""
	if diffSection.file != nil {
		language = diffSection.file.Language
	}

			return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content[1:]))
			return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content[1:]))
		return template.HTML(highlight.Code(diffSection.FileName, language, diffLine.Content))
	diffRecord := diffMatchPatch.DiffMain(highlight.Code(diffSection.FileName, language, diff1[1:]), highlight.Code(diffSection.FileName, language, diff2[1:]), true)
	IsGenerated             bool
	IsVendored              bool
	Language                string
	Start, End                             string

			lastFile := createDiffFile(diff, line)
			diff.End = lastFile.Name
			_, err := io.Copy(io.Discard, reader)
					// The shortest string that can end up here is:
					// "--- a\t\n" without the qoutes.
					// This line has a len() of 7 but doesn't contain a oldName.
					// So the amount that the line need is at least 8 or more.
					// The code will otherwise panic for a out-of-bounds.
					if len(line) > 7 && line[4] == 'a' {
			curSection = &DiffSection{file: curFile}
				curSection = &DiffSection{file: curFile}
				curSection = &DiffSection{file: curFile}
				curSection = &DiffSection{file: curFile}
				count, err := db.Count(m)
		if len(name) == 0 {
			log.Error("Reader has no file name: %v", rd)
			return "", true
		}

func GetDiffRangeWithWhitespaceBehavior(gitRepo *git.Repository, beforeCommitID, afterCommitID, skipTo string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	argsLength := 6
	if len(whitespaceBehavior) > 0 {
		argsLength++
	}
	if len(skipTo) > 0 {
		argsLength++
	}

	diffArgs := make([]string, 0, argsLength)
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
		diffArgs = append(diffArgs, "diff", "--src-prefix=\\a/", "--dst-prefix=\\b/", "-M")
	if skipTo != "" {
		diffArgs = append(diffArgs, "--skip-to="+skipTo)
	}
	cmd := exec.CommandContext(ctx, git.GitExecutable, diffArgs...)

	diff.Start = skipTo

	var checker *git.CheckAttributeReader

	if git.CheckGitVersionAtLeast("1.7.8") == nil {
		indexFilename, worktree, deleteTemporaryFile, err := gitRepo.ReadTreeToTemporaryIndex(afterCommitID)
		if err == nil {
			defer deleteTemporaryFile()

			checker = &git.CheckAttributeReader{
				Attributes: []string{"linguist-vendored", "linguist-generated", "linguist-language", "gitlab-language"},
				Repo:       gitRepo,
				IndexFile:  indexFilename,
				WorkTree:   worktree,
			}
			ctx, cancel := context.WithCancel(git.DefaultContext)
			if err := checker.Init(ctx); err != nil {
				log.Error("Unable to open checker for %s. Error: %v", afterCommitID, err)
			} else {
				go func() {
					err := checker.Run()
					if err != nil && err != ctx.Err() {
						log.Error("Unable to open checker for %s. Error: %v", afterCommitID, err)
					}
					cancel()
				}()
			}
			defer func() {
				cancel()
			}()
		}
	}


		gotVendor := false
		gotGenerated := false
		if checker != nil {
			attrs, err := checker.CheckPath(diffFile.Name)
			if err == nil {
				if vendored, has := attrs["linguist-vendored"]; has {
					if vendored == "set" || vendored == "true" {
						diffFile.IsVendored = true
						gotVendor = true
					} else {
						gotVendor = vendored == "false"
					}
				}
				if generated, has := attrs["linguist-generated"]; has {
					if generated == "set" || generated == "true" {
						diffFile.IsGenerated = true
						gotGenerated = true
					} else {
						gotGenerated = generated == "false"
					}
				}
				if language, has := attrs["linguist-language"]; has && language != "unspecified" && language != "" {
					diffFile.Language = language
				} else if language, has := attrs["gitlab-language"]; has && language != "unspecified" && language != "" {
					diffFile.Language = language
				}
			} else {
				log.Error("Unexpected error: %v", err)
			}
		}

		if !gotVendor {
			diffFile.IsVendored = analyze.IsVendor(diffFile.Name)
		}
		if !gotGenerated {
			diffFile.IsGenerated = analyze.IsGenerated(diffFile.Name)
		}

	separator := "..."
	if directComparison {
		separator = ".."
	}

	shortstatArgs := []string{beforeCommitID + separator + afterCommitID}
func GetDiffCommitWithWhitespaceBehavior(gitRepo *git.Repository, commitID, skipTo string, maxLines, maxLineCharacters, maxFiles int, whitespaceBehavior string, directComparison bool) (*Diff, error) {
	return GetDiffRangeWithWhitespaceBehavior(gitRepo, "", commitID, skipTo, maxLines, maxLineCharacters, maxFiles, whitespaceBehavior, directComparison)