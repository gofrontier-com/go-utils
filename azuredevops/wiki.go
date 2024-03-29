package azuredevops

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/microsoft/azure-devops-go-api/azuredevops/wiki"
)

// CreateWikiIfNotExists creates a code wiki if it does not exist.
func (a *AzureDevOps) CreateWikiIfNotExists(projectName string, wikiName string, gitEmail string, gitUsername string) (*wiki.WikiV2, *git.GitRepository, error) {
	client, err := wiki.NewClient(a.ctx, a.connection)
	if err != nil {
		return nil, nil, err
	}

	getWikiArgs := wiki.GetWikiArgs{
		Project:        &projectName,
		WikiIdentifier: &wikiName,
	}

	r, err := a.createRepositoryIfNotExists(projectName, wikiName, gitEmail, gitUsername)
	if err != nil {
		return nil, nil, err
	}

	w, err := client.GetWiki(a.ctx, getWikiArgs)

	if err == nil {
		return w, r, nil
	}

	// TODO: Check that this is a WikiNotFoundError

	projectId, err := a.getProjectUUID(projectName)
	if err != nil {
		return nil, nil, err
	}

	branch := "main"
	mappedPath := "/"
	wikiCreateArgs := wiki.CreateWikiArgs{
		Project: &projectName,
		WikiCreateParams: &wiki.WikiCreateParametersV2{
			MappedPath:   &mappedPath,
			Name:         &wikiName,
			ProjectId:    projectId,
			Type:         &wiki.WikiTypeValues.CodeWiki,
			RepositoryId: r.Id,
			Version: &git.GitVersionDescriptor{
				VersionType: &git.GitVersionTypeValues.Branch,
				Version:     &branch,
			},
		},
	}
	w, err = client.CreateWiki(a.ctx, wikiCreateArgs)
	if err != nil {
		return nil, nil, err
	}

	return w, r, nil
}
