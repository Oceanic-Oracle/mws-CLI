package profile_storage

type IProfile interface {
	Create(outPath, name, user, project string) error
}