// Code generated by go generate; DO NOT EDIT.
// 2017-12-28 18:55:07.409784145 -0800 PST m=+0.036504731

package locale

var translations = map[string]string{
	"en_US": `{
    "plural.feed.error_count": [
        "%d error",
        "%d errors"
    ],
    "plural.categories.feed_count": [
        "There is %d feed.",
        "There are %d feeds."
    ]
}`,
	"fr_FR": `{
    "plural.feed.error_count": [
        "%d erreur",
        "%d erreurs"
    ],
    "plural.categories.feed_count": [
        "Il y a %d abonnement.",
        "Il y a %d abonnements."
    ],
    "Username": "Nom d'utilisateur",
    "Password": "Mot de passe",
    "Unread": "Non lus",
    "History": "Historique",
    "Feeds": "Abonnements",
    "Categories": "Catégories",
    "Settings": "Réglages",
    "Logout": "Se déconnecter",
    "Next": "Suivant",
    "Previous": "Précédent",
    "New Subscription": "Nouvel Abonnment",
    "Import": "Importation",
    "Export": "Exportation",
    "There is no category. You must have at least one category.": "Il n'y a aucune catégorie. Vous devez avoir au moins une catégorie.",
    "URL": "URL",
    "Category": "Catégorie",
    "Find a subscription": "Trouver un abonnement",
    "Loading...": "Chargement...",
    "Create a category": "Créer une catégorie",
    "There is no category.": "Il n'y a aucune catégorie.",
    "Edit": "Modifier",
    "Remove": "Supprimer",
    "No feed.": "Aucun abonnement.",
    "There is no article in this category.": "Il n'y a aucun article dans cette catégorie.",
    "Original": "Original",
    "Mark this page as read": "Marquer cette page comme lu",
    "not yet": "pas encore",
    "just now": "à l'instant",
    "1 minute ago": "il y a une minute",
    "%d minutes ago": "il y a %d minutes",
    "1 hour ago": "il y a une heure",
    "%d hours ago": "il y a %d heures",
    "yesterday": "hier",
    "%d days ago": "il y a %d jours",
    "%d weeks ago": "il y a %d semaines",
    "%d months ago": "il y a %d mois",
    "%d years ago": "il y a %d ans",
    "Date": "Date",
    "IP Address": "Adresse IP",
    "User Agent": "Navigateur Web",
    "Actions": "Actions",
    "Current session": "Session actuelle",
    "Sessions": "Sessions",
    "Users": "Utilisateurs",
    "Add user": "Ajouter un utilisateur",
    "Choose a Subscription": "Choisissez un abonnement",
    "Subscribe": "S'abonner",
    "New Category": "Nouvelle Catégorie",
    "Title": "Titre",
    "Save": "Sauvegarder",
    "or": "ou",
    "cancel": "annuler",
    "New User": "Nouvel Utilisateur",
    "Confirmation": "Confirmation",
    "Administrator": "Administrateur",
    "Edit Category: %s": "Modification de la catégorie : %s",
    "Update": "Mettre à jour",
    "Edit Feed: %s": "Modification de l'abonnement : %s",
    "There is no category!": "Il n'y a aucune catégorie !",
    "Edit user: %s": "Modification de l'utilisateur : %s",
    "There is no article for this feed.": "Il n'y a aucun article pour cet abonnement.",
    "Add subscription": "Ajouter un abonnement",
    "You don't have any subscription.": "Vous n'avez aucun abonnement",
    "Last check:": "Dernière vérification :",
    "Refresh": "Actualiser",
    "There is no history at the moment.": "Il n'y a aucun historique pour le moment.",
    "OPML file": "Fichier OPML",
    "Sign In": "Connexion",
    "Sign in": "Connexion",
    "Theme": "Thème",
    "Timezone": "Fuseau horaire",
    "Language": "Langue",
    "There is no unread article.": "Il n'y a rien de nouveau à lire.",
    "You are the only user.": "Vous êtes le seul utilisateur.",
    "Last Login": "Dernière connexion",
    "Yes": "Oui",
    "No": "Non",
    "This feed already exists (%s).": "Cet abonnement existe déjà (%s).",
    "Unable to fetch feed (statusCode=%d).": "Impossible de récupérer cet abonnement (code=%d).",
    "Unable to open this link: %v": "Impossible d'ouvrir ce lien : %v",
    "Unable to analyze this page: %v": "Impossible d'analyzer cette page : %v",
    "Unable to find any subscription.": "Impossible de trouver un abonnement.",
    "The URL and the category are mandatory.": "L'URL et la catégorie sont obligatoire.",
    "All fields are mandatory.": "Tous les champs sont obligatoire.",
    "Passwords are not the same.": "Les mots de passe ne sont pas les mêmes.",
    "You must use at least 6 characters.": "Vous devez utiliser au moins 6 caractères.",
    "The username is mandatory.": "Le nom d'utilisateur est obligatoire.",
    "The username, theme, language and timezone fields are mandatory.": "Le nom d'utilisateur, le thème, la langue et le fuseau horaire sont obligatoire.",
    "The title is mandatory.": "Le titre est obligatoire.",
    "About": "A propos",
    "version": "Version",
    "Version:": "Version :",
    "Build Date:": "Date de la compilation :",
    "Author:": "Auteur :",
    "Authors": "Auteurs",
    "License:": "Licence :",
    "Attachments": "Pièces jointes",
    "Download": "Télécharger",
    "Invalid username or password.": "Mauvais identifiant ou mot de passe.",
    "Never": "Jamais",
    "Unable to execute request: %v": "Impossible d'exécuter cette requête: %v",
    "Last Parsing Error": "Dernière erreur d'analyse",
    "There is a problem with this feed": "Il y a un problème avec cet abonnement",
    "Unable to parse OPML file: %v.": "Impossible de lire ce fichier OPML : %v.",
    "Unable to parse RSS feed: %v.": "Impossible de lire ce flux RSS: %v.",
    "Unable to parse Atom feed: %v.": "Impossible de lire ce flux Atom: %v.",
    "Unable to parse JSON feed: %v.": "Impossible de lire ce flux JSON: %v.",
    "Unable to parse RDF feed: %v.": "Impossible de lire ce flux RDF: %v.",
    "Unable to normalize encoding: %v.": "Impossible de normaliser l'encodage : %v.",
    "Unable to create this category.": "Impossible de créer cette catégorie.",
    "yes": "oui",
    "no": "non",
    "Are you sure?": "Êtes-vous sûr ?",
    "Work in progress...": "Travail en cours...",
    "This user already exists.": "Cet utilisateur existe déjà.",
    "This category already exists.": "Cette catégorie existe déjà.",
    "Unable to update this category.": "Impossible de mettre à jour cette catégorie.",
    "Integrations": "Intégrations",
    "Bookmarklet": "Bookmarklet",
    "Drag and drop this link to your bookmarks.": "Glisser-déposer ce lien dans vos favoris.",
    "This special link allows you to subscribe to a website directly by using a bookmark in your web browser.": "Ce lien spécial vous permet de vous abonner à un site web directement en utilisant un marque page dans votre navigateur web.",
    "Add to Miniflux": "Ajouter à Miniflux",
    "Refresh all feeds in background": "Actualiser les abonnements en arrière-plan",
    "Sign in with Google": "Se connecter avec Google",
    "Unlink my Google account": "Dissocier mon compte Google",
    "Link my Google account": "Associer mon compte Google",
    "Category not found for this user.": "Cette catégorie n'existe pas pour cet utilisateur.",
    "Invalid theme.": "Le thème est invalide.",
    "Entry Sorting": "Ordre des éléments",
    "Older entries first": "Ancien éléments en premier",
    "Recent entries first": "Éléments récents en premier",
    "Saving...": "Enregistrement...",
    "Done!": "Terminé !",
    "Save this article": "Sauvegarder cet article",
    "Mark bookmark as unread": "Marquer le lien comme non lu",
    "Pinboard Tags": "Libellés de Pinboard",
    "Pinboard API Token": "Jeton de sécurité de l'API de Pinboard",
    "Save articles to Pinboard": "Sauvegarder les articles vers Pinboard",
    "Save articles to Instapaper": "Sauvegarder les articles vers Instapaper",
    "Instapaper Username": "Nom d'utilisateur Instapaper",
    "Instapaper Password": "Mot de passe Instapaper",
    "Activate Fever API": "Activer l'API de Fever",
    "Fever Username": "Nom d'utilisateur pour l'API de Fever",
    "Fever Password": "Mot de passe pour l'API de Fever",
    "Fetch original content": "Récupérer le contenu original",
    "Scraper Rules": "Règles pour récupérer le contenu original",
    "Rewrite Rules": "Règles de réécriture",
    "Preferences saved!": "Préférences sauvegardées !",
    "Your external account is now linked !": "Votre compte externe est maintenant associé !",
    "Save articles to Wallabag": "Sauvegarder les articles vers Wallabag",
    "Wallabag API Endpoint": "URL de l'API de Wallabag",
    "Wallabag Client ID": "Identifiant du client Wallabag",
    "Wallabag Client Secret": "Clé secrète du client Wallabag",
    "Wallabag Username": "Nom d'utilisateur de Wallabag",
    "Wallabag Password": "Mot de passe de Wallabag",
    "Keyboard Shortcut: %s": "Raccourci clavier : %s",
    "Favorites": "Favoris",
    "Star": "Favoris",
    "Unstar": "Enlever favoris",
    "Starred": "Favoris",
    "There is no bookmark at the moment.": "Il n'y a aucun favoris pour le moment.",
    "Last checked:": "Dernière vérification :",
    "ETag header:": "En-tête ETag :",
    "LastModified header:": "En-tête LastModified :",
    "None": "Aucun/Aucune",
    "Keyboard Shortcuts": "Raccourcis clavier",
    "Sections Navigation": "Naviguation entre les sections",
    "Go to unread": "Aller aux éléments non lus",
    "Go to bookmarks": "Aller aux favoris",
    "Go to history": "Voir l'historique",
    "Go to feeds": "Voir les abonnements",
    "Go to categories": "Voir les catégories",
    "Go to settings": "Voir les réglages",
    "Show keyboard shortcuts": "Voir les raccourcis clavier",
    "Items Navigation": "Naviguation entre les éléments",
    "Go to previous item": "Élément précédent",
    "Go to next item": "Élément suivant",
    "Go to previous page": "Page précédente",
    "Go to next page": "Page suivante",
    "Open selected item": "Ouvrir élément sélectionné",
    "Open original link": "Ouvrir lien original",
    "Toggle read/unread": "Basculer entre lu/non lu",
    "Mark current page as read": "Marquer la page actuelle comme lu",
    "Download original content": "Télécharger le contenu original",
    "Toggle bookmark": "Ajouter/Enlever favoris",
    "Close modal dialog": "Fermer la boite de dialogue",
    "Save article": "Sauvegarder l'article"
}
`,
}

var translationsChecksums = map[string]string{
	"en_US": "6fe95384260941e8a5a3c695a655a932e0a8a6a572c1e45cb2b1ae8baa01b897",
	"fr_FR": "30f70cf369dae3e0461e44a444be56d657d7d381801c321e7312886e75278c81",
}
