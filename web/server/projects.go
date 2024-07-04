package server

type Project struct {
	Name     string
	ImgSrc   string
	RepoLink string
	DemoLink string
	Tools    []string
}

var MyProjects = []Project{
	{
		Name:     "Subway PathFinder",
		ImgSrc:   "subwaypathfinder.png",
		RepoLink: "https://github.com/VitorCarraraMarques/subway",
		Tools:    []string{"python.svg", "fastapi.svg", "html.svg", "javascript.svg", "ajax.svg"},
	},
	{
		Name:     "LandingPageV1",
		ImgSrc:   "landingpagev1.png",
		RepoLink: "https://github.com/VitorCarraraMarques/LandingPage",
		Tools:    []string{"python.svg", "django.svg", "html.svg", "css.svg"},
	},
	{
		Name:     "Planejador de Móveis",
		ImgSrc:   "movelplanner.png",
		RepoLink: "https://github.com/VitorCarraraMarques/organizador-de-moveis",
		DemoLink: "https://vitorcarraramarques.github.io/organizador-de-moveis/",
		Tools:    []string{"html.svg", "css.svg", "javascript.svg", "p5js.svg"},
	},
	{
		Name:     "Ice Skater Game",
		ImgSrc:   "skatergame.png",
		RepoLink: "https://github.com/VitorCarraraMarques/Ice-Skater-Game",
		Tools:    []string{"python.svg", "pygame.svg"},
	},
	{
		Name:     "Somador de Ondas",
		ImgSrc:   "ondasoma.png",
		DemoLink: "https://editor.p5js.org/vitorcarrara/sketches/TkHiotTZx",
		Tools:    []string{"javascript.svg", "p5js.svg"},
	},
	{
		Name:     "NEWTONIÆN PHYSICA SIMULÆTOR",
		ImgSrc:   "simphys.png",
		RepoLink: "https://github.com/VitorCarraraMarques/SimuladorFisicaPygame",
		DemoLink: "https://editor.p5js.org/vitorcarrara/full/QQ8uXPbqP",
		Tools:    []string{"javascript.svg", "p5js.svg", "python.svg", "pygame.svg"},
	},
	{
		Name:     "éPrimo?!",
		ImgSrc:   "ehprimo.png",
		RepoLink: "https://github.com/VitorCarraraMarques/AplicativoEhPrimo",
		Tools:    []string{"python.svg", "kivy.svg"},
	},
}
