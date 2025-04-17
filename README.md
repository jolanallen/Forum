Bien sûr, voici le `README.md` mis à jour pour inclure l'installation d'OpenSSL comme prérequis :

<<<<<<< HEAD
# Projet Ynov
=======
```markdown
# Projet Ynov

## Forum Application

Cette application est un forum simple basé sur Go et Docker, avec une base de données MySQL et un serveur sécurisé en HTTPS. Ce projet inclut des conteneurs Docker pour l'application et la base de données.

## Prérequis

Avant de commencer, assurez-vous d'avoir installé les outils suivants sur votre machine :

- **Docker** : [Installation de Docker](https://docs.docker.com/get-docker/)
- **Docker Compose** : [Installation de Docker Compose](https://docs.docker.com/compose/install/)
- **Go** : [Installation de Go](https://golang.org/doc/install/)
- **OpenSSL** : [Installation d'OpenSSL](https://www.openssl.org/)

## Installation

1. **Clonez ce projet** sur votre machine :

   ```bash
   git clone https://votre-repository.git
   cd Forum
   ```

2. **Exécutez le script `launcher.sh`** :

   Ce script va vérifier les prérequis, générer un certificat SSL autosigné si nécessaire, construire les images Docker pour l'application et la base de données, et lancer les conteneurs nécessaires.

   ```bash
   ./launcher.sh
   ```

   Le script réalisera les actions suivantes :
   - Vérification des outils nécessaires (Docker, Docker Compose, Go, OpenSSL).
   - Génération des certificats SSL/TLS si ces derniers n'existent pas déjà (à l'aide d'OpenSSL).
   - Construction des images Docker pour l'application (`forum-app`) et la base de données (`forum-mysql`).
   - Démarrage des conteneurs Docker via `docker-compose`.

## Accès à l'application

Une fois les conteneurs démarrés, vous pouvez accéder à l'application depuis votre réseau domestique.

- Si vous travaillez sur la machine hôte, ouvrez votre navigateur à l'adresse suivante :
  ```
  https://localhost/forum/
  ```
  
- Depuis un autre appareil connecté à votre réseau domestique, remplacez `<your_ip>` par l'adresse IP de la machine hôte :
  ```
  https://<your_ip>/forum/
  ```

## Structure des répertoires

Voici la structure des fichiers du projet :

```
.
├── arborescence.txt
├── backend
│   ├── db
│   │   ├── db.go
│   │   └── queries.go
│   ├── handler
│   │   ├── admin.go
│   │   ├── authentification.go
│   │   ├── guest.go
│   │   └── user.go
│   ├── middlewares
│   │   ├── auth.go
│   │   ├── cookie.go
│   │   ├── logger.go
│   │   └── security.go
│   ├── server
│   │   ├── routes.go
│   │   ├── server.go
│   │   └── ssl_tls
│   │       ├── cert.pem
│   │       └── key.pem
│   ├── session
│   │   ├── session.go
│   │   └── session_store.go
│   ├── structs
│   │   ├── admin.go
│   │   ├── comment.go
│   │   ├── Forum.go
│   │   ├── guest.go
│   │   ├── image.go
│   │   ├── post.go
│   │   └── user.go
│   ├── UnitTest
│   │   ├── Unit_Test_Auth.go
│   │   ├── Unit_Test_DB_Queries.go
│   │   └── Unit_Test_Logic.go
│   └── utils
│       ├── email.go
│       ├── Init.go
│       ├── password.go
│       ├── topic.go
│       └── validations.go
├── databases
│   ├── Dockerfile
│   └── forum.sql
├── docker-compose.yaml
├── Dockerfile
├── forum-app
├── forum image sql
│   ├── Capture d'écran 2025-03-10 104158.png
│   ├── Capture d'écran 2025-03-10 232426.png
│   └── creation d'un user pour exporter.png
├── go.mod
├── go.sum
├── launcher.sh
├── main
│   └── main.go
├── README.md
├── route management.jpg
├── structurefichier.txt
└── web
    ├── static
    │   ├── css
    │   │   ├── index.css
    │   │   ├── newPost.css
    │   │   ├── profil.css
    │   │   └── waitingRoom.css
    │   ├── img
    │   │   ├── logo.2023a.png
    │   │   └── logo.jpg
    │   └── js
    │       ├── index.js
    │       ├── post.js
    │       └── waitingRoom.js
    └── templates
        ├── admin
        │   ├── dashboard.html
        │   ├── manage_post.html
        │   └── manage_user.html
        ├── auth
        │   ├── login.html
        │   ├── logout.html
        │   └── register.html
        ├── errors
        │   ├── 403.html
        │   ├── 404.html
        │   └── 500.html
        ├── forum
        │   ├── home.html
        │   ├── new_post.html
        │   ├── sujet.html
        │   └── waitingRoom.html
        ├── include
        │   ├── comment_card.html
        │   ├── footer.html
        │   ├── header.html
        │   ├── navbar.html
        │   └── post_card.html
        ├── index.html
        └── users
            ├── profile.html
            └── profile_settings.html

26 directories, 77 files

```

## Variables d'environnement

Le serveur utilise les variables suivantes pour la configuration des certificats SSL/TLS :

- `CERT_PATH` : Le chemin vers le fichier de certificat SSL.
- `KEY_PATH` : Le chemin vers le fichier de clé privée SSL.

Ces variables sont automatiquement définies dans le script `launcher.sh` et dans le fichier `.env` pour Docker.

## Résolution des problèmes

Si vous rencontrez des erreurs liées aux certificats SSL, assurez-vous que les fichiers de certificat sont présents dans le répertoire `backend/server/ssl_tls/` et que les variables d'environnement **`CERT_PATH`** et **`KEY_PATH`** pointent correctement vers ces fichiers.

Si le script échoue à générer les certificats SSL, assurez-vous que **OpenSSL** est installé et disponible dans votre environnement. Vous pouvez vérifier la présence d'OpenSSL avec la commande suivante :

```bash
openssl version
```

## Contribuer

1. Fork ce repository.
2. Créez une branche pour votre fonctionnalité (`git checkout -b feature/nom-de-la-fonctionnalité`).
3. Effectuez vos modifications.
4. Faites un commit avec un message clair (`git commit -am 'Ajoute la fonctionnalité X'`).
5. Poussez vos modifications (`git push origin feature/nom-de-la-fonctionnalité`).
6. Ouvrez une Pull Request.

## License

Ce projet est sous la licence MIT. Consultez le fichier [LICENSE](LICENSE) pour plus de détails.
```

### Modifications apportées :
- Ajout de **OpenSSL** dans la section **Prérequis** avec un lien vers la documentation pour l'installation.
- Mention d'OpenSSL dans la section **Résolution des problèmes**, avec une commande pour vérifier qu'il est installé correctement.

Cela devrait aider les utilisateurs à préparer correctement leur environnement avant de commencer avec le projet.
>>>>>>> af63cc477bd8d8c040f360bc1a17fa764cd57e4d
