# CaptainFeedHook

CaptainFeedHook est une application de gestion de flux RSS qui publie automatiquement des mises à jour provenant de différentes sources RSS vers des Webhooks Discord. Elle est idéale pour garder vos communautés informées des dernières actualités, vidéos ou articles.

## Fonctionnalités

- **Support multi-sources** : Configurez plusieurs flux RSS avec leurs propres Webhooks Discord.
- **Intervalles de mise à jour personnalisables** : Définissez la fréquence de mise à jour pour chaque flux RSS.
- **Facilité de configuration** : Utilise un fichier `config.json` simple pour gérer les flux et les Webhooks.
- **Léger et rapide** : Construit avec Go pour une performance optimale.

## Prérequis

- Go SDK ou Docker installé sur votre machine.
- JSON valide pour le fichier de configuration `config.json`.

## Configuration des flux 

Dans le fichier `config.json`, ajoutez une entrée pour chaque flux RSS. Exemple :

```json
{
  "example": {
    "rss": "https://www.example.com/.rss",
    "webhook": "https://discord.com/api/webhooks/123456789/example",
    "interval": 600
  }
}
```

| Champ        | Description                                                                 |
|--------------|-----------------------------------------------------------------------------|
| `rss`        | URL du flux RSS.                                                           |
| `webhook`    | URL du Webhook Discord pour publier les mises à jour.                      |
| `interval`   | Intervalle en secondes entre les vérifications des mises à jour.           |

## Utilisation (Go)

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/AquaBx/CaptainFeedHook.git
   cd CaptainFeedHook
   ```

2. Créer un dossier 'config' et créez un fichier `config.json` à l'intérieur :

    Voir [Configuration des flux](#configuration-des-flux) pour créer ce fichier.

3. Compiler l'application
    ```bash
    go build
   ```

4. Lancez l'application :
   ```bash
    ./CaptainFeedHook
   ```

## Utilisation (Docker)

1. Créez un fichier `config.json` dans votre répertoire :

    Voir [Configuration des flux](#configuration-des-flux) pour créer ce fichier.

2. Créez un fichier Dockerfile :
   ```Dockerfile
   FROM ghcr.io/aquabx/captainfeedhook:latest
   COPY config.json /config/config.json
   ```

3. Construisez l'image Docker :
   ```bash
   docker build -t captainfeedhook .
   ```

4. Lancez le conteneur :
   ```bash
   docker run -d --name captainfeedhook captainfeedhook
   ```


## Contribution

Les contributions sont les bienvenues ! Si vous souhaitez ajouter une fonctionnalité ou corriger un bug, n'hésitez pas à ouvrir une issue ou une pull request.

## Licence

Ce projet est sous licence GNU GPL. Consultez le fichier [LICENSE](LICENSE) pour plus de détails.