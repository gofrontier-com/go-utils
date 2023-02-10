package git

type Git interface {
	Add(path string) error
	Checkout(branchName string, create bool) error
	CloneOverHttp(url string, username string, password string) error
	Commit(message string) error
	GetRepositoryPath() string
	NewGit(repositoryPath string) *Git
	Push(force bool) error
}