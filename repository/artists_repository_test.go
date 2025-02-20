package repository

func (s *DbTestSuite) TestFetchSuccessfully() {
	artist, err := FetchAll(NewArtistsRepository(&s.db))
	if err != nil {
		panic(err)
	}
	s.NotNil(artist)
}

func (s *DbTestSuite) TestFetchByIdSuccessfully() {
	artist, err := FetchById(NewArtistsRepository(&s.db), 1)

	if err != nil {
		panic(err)
	}

	s.NotNil(artist)

}

func (s *DbTestSuite) TestMockAlteracoesModuloTerraform() {
	terraform, err := FetchById(NewArtistsRepository(&s.db), 1)
	s.NotNil(terraform)

}
