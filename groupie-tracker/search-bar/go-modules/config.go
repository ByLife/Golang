package Groupie

import "html/template"

/*

FR:

Client d'une API pour https://www.groupietrackers.herokuapp.com

Vous pouvez le modifier, aucun soucis !
S'il y a un problème quelconque:
-> EMAIL: bylifeh@gmail.com
-> DISCORD: ByLife#2988

Tout le code est facile à lire et normallement,
Il est facile à modifier pour votre client d'API (Je le recommande pas car cette API peut s'utiliser d'une autre manière que la votre donc faudra tout recoder et beaucoup modifier)
Si c'est juste pour cette API, pas de problème !

/!\Si tu as une erreur par rapport à un répertoire non trouvé, il faut bien que tu sois dans le bon répertoire/!\
Dans tous les cas, si besoin, vous pouvez aussi modifer les configurations et les répertoires ci-dessous après ce commentaire.

Bonne journée/soirée à toi ! :)

EN:

API CLIENT FOR https://www.groupietrackers.herokuapp.com

CAN BE MODIFIED, ALL RIGHTS TO EVERYONE !
ANY PROBLEM, MY CONTACTS ARE:
-> EMAIL: bylifeh@gmail.com
-> DISCORD: ByLife#2988

All the code is readable and normally,
Easy to modify if needed for your client (I don't recommand cuz used for this API so can be really different from yours)
If it's only for this API, no problem !

/!\If getting error about a file directory or file name not found, careful of the directory that you are in/!\
If needed, you can also change the files directories  bellow.

HOPE YOU HAVE A GREAT DAY !

*/

var (
	//API Base Link && Web Hosting Port

	API_HTTP_OR_HTTPS = "https://"
	API_BaseLink      = "groupietrackers.herokuapp.com"

	Web_PORT               = ":8080"
	Web_AssetsDirectory    = "assets"
	Web_CssDirectory       = Web_AssetsDirectory + "/css"
	Web_TemplatesDirectory = Web_AssetsDirectory + "/templates/*.html"

	// API Directories

	API_Directory         = "/api"
	API_ArtistDirectory   = "/artists"
	API_LocationDirectory = "/location"
	API_RelationDirectory = "/relation"
	API_DateDirectory     = "/dates"

	// Addition of everything so final result will be https://xxx.yyy.com/api_base_link/api_directories

	API_LinkToArtist   = API_HTTP_OR_HTTPS + API_BaseLink + API_Directory + API_ArtistDirectory
	API_LinkToLocation = API_HTTP_OR_HTTPS + API_BaseLink + API_Directory + API_LocationDirectory
	API_LinkToRelation = API_HTTP_OR_HTTPS + API_BaseLink + API_Directory + API_RelationDirectory
	API_LinkToDates    = API_HTTP_OR_HTTPS + API_BaseLink + API_Directory + API_DateDirectory

	// Loading Template Handler Variable

	WebHtml     *template.Template
	ApiData     Artist
	ApiDataTemp Artist
	SearchValue string
)
