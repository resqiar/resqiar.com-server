package constants

type stack struct {
	Title string
	URL   string
}

var svelte = stack{
	Title: "Svelte",
	URL:   "/media/techs/svelte.webp",
}

var typescript = stack{
	Title: "TypeScript",
	URL:   "/media/techs/ts.webp",
}

var astro = stack{
	Title: "Astro",
	URL:   "/media/techs/astro.webp",
}

var tailwind = stack{
	Title: "Tailwind CSS",
	URL:   "/media/techs/tailwind.webp",
}

var jdoodle = stack{
	Title: "Jdoodle API",
	URL:   "/media/techs/jdoodle.webp",
}

var nestjs = stack{
	Title: "NestJS",
	URL:   "/media/techs/nestjs.webp",
}

var fastify = stack{
	Title: "Fastify",
	URL:   "/media/techs/fastify.webp",
}

var postgres = stack{
	Title: "PostgreSQL",
	URL:   "/media/techs/postgres.webp",
}

var typeorm = stack{
	Title: "TypeORM",
	URL:   "/media/techs/typeorm.webp",
}

var passport = stack{
	Title: "PassportJS",
	URL:   "/media/techs/passport.webp",
}

var python = stack{
	Title: "Python",
	URL:   "/media/techs/python.webp",
}

var flask = stack{
	Title: "Flask",
	URL:   "/media/techs/flask.webp",
}

var golang = stack{
	Title: "Golang",
	URL:   "/media/techs/golang.webp",
}

var ShowcaseData = []struct {
	Title       string
	Description string
	DemoURL     string
	Images      []string
	SourceURL   string
	Techs       []stack
}{
	{
		Title:       "Resqiar.com",
		Description: "My personal page where it becomes the place I want to share all the knowledges, thoughts, playground and everything in between.",
		Images: []string{
			"/media/resqiar-1.webp",
			"/media/resqiar-2.webp",
			"/media/resqiar-3.webp",
			"/media/resqiar-4.webp",
			"/media/resqiar-5.webp",
			"/media/resqiar-6.webp",
		},
		DemoURL:   "https://resqiar.com",
		SourceURL: "https://github.com/resqiar/resdev",
		Techs:     []stack{svelte, typescript, tailwind, golang, postgres},
	},
	{
		Title:       "Roof Tile Collection",
		Description: "A collections of Roof Tile for SME Business",
		Images:      []string{"/media/rooftile-1.webp", "/media/rooftile-2.webp"},
		DemoURL:     "https://tokoriskiageng.vercel.app",
		Techs:       []stack{astro, tailwind},
	},
	{
		Title:       "Algo Visualizer",
		Description: "Visualize infamous many algorithms for showcase and learning purpose.",
		Images:      []string{"/media/algo.webp", "/media/algo-1.webp", "/media/algo-2.webp", "/media/algo-3.webp"},
		DemoURL:     "https://resqiar.github.io/algo-visualization",
		SourceURL:   "https://github.com/resqiar/algo-visualization",
		Techs:       []stack{svelte, typescript},
	},
	{
		Title:       "Binder",
		Description: "My personal binder extensions. Used to store additional data to organize my binder books. With Image, Code, Playground, and QR Code functionalities, effortlessly store and organize additional data within your binder books.",
		Images: []string{
			"/media/binder.webp",
			"/media/binder-1.webp",
			"/media/binder-2.webp",
			"/media/binder-3.webp",
		},
		SourceURL: "https://github.com/resqiar/binder",
		Techs:     []stack{svelte, typescript, tailwind, jdoodle},
	},
	{
		Title:       "Binder Server",
		Description: "Binder server built with NestJS and Fastify. The server provides and maintains Extension's data to Binder App through its APIs.",
		Images: []string{
			"/media/binder.webp",
			"/media/binder-1.webp",
			"/media/binder-2.webp",
			"/media/binder-3.webp",
		},
		SourceURL: "https://github.com/resqiar/binder-server",
		Techs:     []stack{typescript, nestjs, fastify, postgres, typeorm, passport},
	},
	{
		Title:       "Go Bookstore",
		Description: "Go Bookstore is a Golang and Postgres-based project that offers an online book store. The project also includes an admin dashboard for management of inventory. in collaboration with @Hilll19 and @nathanpasca",
		Images: []string{
			"/media/bookstore.webp",
			"/media/bookstore-1.webp",
			"/media/bookstore-2.webp",
			"/media/bookstore-3.webp",
		},
		SourceURL: "https://github.com/resqiar/bookstore",
		Techs:     []stack{golang, postgres},
	},
	{
		Title:       "AI Anime Recommender",
		Description: "Personalized anime recommendation based on user-based similarity calculations. Use collaborative filtering method with Pearson and Cosine. Built in-collaboration with @Hilll19 and @nathanpasca.",
		Images:      []string{"/media/ai-recommender.webp"},
		SourceURL:   "https://github.com/resqiar/anime-recommender",
		Techs:       []stack{python, flask},
	},
}
