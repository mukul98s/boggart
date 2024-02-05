package data

type VersionLink struct {
	DownloadLink string
	Sha256Sum    string
}

var Versions = map[string]VersionLink{
	"7.0": {
		DownloadLink: "https://www.php.net/distributions/php-7.0.33.tar.gz",
		Sha256Sum:    "d71a6ecb6b13dc53fed7532a7f8f949c4044806f067502f8fb6f9facbb40452a",
	},
	"7.3": {
		DownloadLink: "https://www.php.net/distributions/php-7.3.33.tar.gz",
		Sha256Sum:    "9a369c32c6f52036b0a890f290327f148a1904ee66aa56e2c9a7546da6525ec8",
	},
	"8.1": {
		DownloadLink: "https://www.php.net/distributions/php-8.1.27.tar.gz",
		Sha256Sum:    "9aa5d7a29389d799885d90740932697006d5d0f55d1def67678e0c14f6ab7b2d",
	},
	"8.2": {
		DownloadLink: "https://www.php.net/distributions/php-8.2.14.tar.gz",
		Sha256Sum:    "4c1fbb55a10ece7f4532feba9f3f88b9b211c11320742977588738374c03255f",
	},
	"8.3": {
		DownloadLink: "https://www.php.net/distributions/php-8.3.1.tar.gz",
		Sha256Sum:    "2b10218b5e81915d1708ab4b6351362d073556ec73a790553c61fd89c119924e",
	},
}
