#!/bin/bash

# Vérification des prérequis
echo "🔍 Vérification des outils nécessaires..."

# Vérifier Docker
if ! command -v docker &> /dev/null; then
    echo "❌ Docker n'est pas installé. Installez Docker et relancez le script."
    exit 1
fi

# Vérifier Docker Compose
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose n'est pas installé. Installez-le et relancez le script."
    exit 1
fi

# Vérifier Go
if ! command -v go &> /dev/null; then
    echo "❌ Go n'est pas installé. Installez Go et relancez le script."
    exit 1
fi

echo "✅ Tous les outils nécessaires sont installés."


# Nettoyer les dépendances et compiler
echo "📦 Mise à jour des dépendances..."
go mod tidy || { echo "❌ Erreur lors de la mise à jour des dépendances Go"; exit 1; }

echo "⚙️ Compilation de l'application..."
go build -o forum-app main/main.go || { echo "❌ Échec de la compilation"; exit 1; }

# Vérification de l'existence des certificats
cert_dir="backend/server/ssl_tls"
cert_file="$cert_dir/cert.pem"
key_file="$cert_dir/key.pem"

if [[ -f "$cert_file" && -f "$key_file" ]]; then
    echo "✅ Les certificats existent déjà, pas besoin de les recréer."
else
    echo "🔑 Génération du certificat autosigné..."

    # Création du répertoire pour les certificats si nécessaire
    mkdir -p "$cert_dir"

    # Génération des certificats auto-signés
    openssl req -x509 -newkey rsa:4096 -keyout "$key_file" -out "$cert_file" -days 365 -nodes -subj "/C=FR/ST=France/L=Paris/O=ForumApp/CN=localhost"

    # Vérification de la génération des certificats
    if [ ! -f "$cert_file" ] || [ ! -f "$key_file" ]; then
        echo "❌ Échec de la génération des certificats autosignés"
        exit 1
    fi

    echo "✅ Certificats générés avec succès !"
fi

# Export des variables d'environnement
echo "🌍 Exportation des variables d'environnement pour les certificats..."
export CERT_PATH=$cert_file
export KEY_PATH=$key_file

echo "✅ Variables d'environnement exportées : CERT_PATH=$CERT_PATH, KEY_PATH=$KEY_PATH"


# Construire les images Docker
echo "🐳 Construction de l'image Docker de l'application..."
docker build -t forum-app . || { echo "❌ Échec de la construction de l'image forum-app"; exit 1; }

echo "🐳 Construction de l'image Docker de la base de données..."
cd databases || { echo "❌ Impossible d'accéder au dossier databases"; exit 1; }
docker build -t forum-mysql . || { echo "❌ Échec de la construction de l'image forum-mysql"; exit 1; }
cd ..

# Lancer les conteneurs
echo "🚀 Démarrage des conteneurs..."
docker-compose up -d || { echo "❌ Échec du démarrage des conteneurs"; exit 1; }

echo "✅ Projet lancé avec succès !"
if [ -d "main" ]; then
    cd main || { echo "❌ Impossible d'entrer dans le répertoire main"; exit 1; }
fi
go run .  # Ou une commande Go spécifique si nécessaire

echo "https://localhost:443/forum/"
