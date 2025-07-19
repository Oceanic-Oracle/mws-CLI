package profile_storage

type IProfile interface {
	Create(outPath, name, user, project string) error
	Delete(outPath, name string) error
	Get(outPath, name string) (string, error)
}