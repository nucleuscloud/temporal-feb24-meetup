package datagen

type Activities struct{}

func (a *Activities) GetGenerateConfig() (string, error) {
	return ``, nil
}

func (a *Activities) SynchronizeTable(
	config string,
) error {

	return nil
}
