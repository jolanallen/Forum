# Étape 1 : Utiliser une image Go pour construire le projet
FROM golang:1.24-alpine as builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers de modules Go dans le répertoire de travail du conteneur
COPY go.mod go.sum ./

# Télécharger les dépendances Go
RUN go mod tidy

# Copier tout le code source du dossier 'main' dans le conteneur
COPY . .

# Se déplacer dans le dossier 'main' et compiler le projet
WORKDIR /app/main

RUN go build -o main .

# Étape 2 : Créer une image finale plus légère basée sur Alpine
FROM alpine:latest

# Installer les bibliothèques nécessaires (si besoin pour l'exécution)
RUN apk --no-cache add ca-certificates

# Copier le binaire compilé depuis l'étape précédente
COPY --from=builder /app/main /usr/local/bin/main

# Définir le binaire comme point d'entrée du conteneur
ENTRYPOINT ["/usr/local/bin/main"]

# Commande par défaut (optionnelle)
CMD []
