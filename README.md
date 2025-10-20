# Puissance 4 (Power 4)

Un jeu de Puissance 4 (Connect Four) en Go avec une interface web responsive, plusieurs niveaux de difficulté et un mode « gravité » optionnel qui inverse la chute des pions toutes les 5 actions.

## ✨ Fonctionnalités

- 3 niveaux de difficulté avec des grilles adaptées:
  - Facile: 6 × 7
  - Normal: 6 × 9
  - Difficile: 7 × 8
- Interface responsive avec grille et indicateurs de colonnes.
- Détection de victoire et d’égalité, bannière de résultat, bouton « Recommencer ou nouvelle partie ».
- Mode Gravité (optionnel):
  - Quand activé, la gravité bascule toutes les 5 poses de pion.
  - Gravité normale: les pions tombent de haut en bas.
  - Gravité inversée: les pions « montent » du bas vers le haut.
  - Un indicateur visuel et un changement de couleur de plateau signalent l’état courant.

## 🛠️ Stack technique

- Langage: Go (>= 1.22)
- Serveur: net/http
- Templates: html/template (pages et partials)
- Styles: CSS (static/style.css, static/play.css)
- Aucune base de données requise

## 📁 Structure du projet

```
.
├── go.mod
├── main.go
├── index.html              # Page d’accueil (choix de difficulté, noms des joueurs, option gravité)
├── pages/
│   └── play.html           # Vue principale du jeu
├── src/
│   ├── game.go             # Logique du jeu et handlers
│   ├── handler.go          # Handler de la page d’accueil
│   ├── serveur.go          # Routage et serveur HTTP
│   └── utils.go            # Tailles de grille et vérification de victoire
├── static/
│   ├── style.css           # Styles de la page d’accueil
│   └── play.css            # Styles de la page de jeu
└── templates/
    ├── board.html          # Template partiel pour la grille
    └── result.html         # Template partiel pour la bannière de résultat
```

## 🚀 Lancer en local

Prérequis: Go 1.22 ou plus récent.

1) Lancer le serveur depuis la racine du projet:

```bash
# Depuis la racine du projet
go run main.go
```

2) Ouvrez dans votre navigateur:

- http://localhost:8080


## 🎮 Jouer

1) Sur la page d’accueil:
   - Choisissez un niveau de difficulté.
   - Entrez les noms des joueurs .
   - Cochez « Activer mode gravité » si vous voulez la gravité .
2) Cliquez sur « Commencer ».
3) Dans la page de jeu, cliquez sur les indicateurs de colonne pour déposer un pion.
4) Fin de partie: une bannière indique le gagnant ou l’égalité et propose « Recommencer ou nouvelle partie ».

### Mode Gravité (optionnel)

- Quand activé, toutes les 5 poses de pions validées, la gravité bascule.
- Un Pop-up apparait indiquant:
  - Gravité normale: chute haut → bas.
  - Gravité inversée: chute bas → haut.

## 🔌 Routes principales

- GET `/` — Page d’accueil.
- POST `/play` — Démarre une partie selon le formulaire d’accueil (difficulté, joueurs, option gravité).
- POST `/move` — Joue un coup en colonne (`col` dans le formulaire). Ignoré si la partie est terminée.
- POST `/rematch` — Relance une partie avec les mêmes paramètres.
- Static `/static/*` — Fichiers statiques (CSS).

## ⚙️ Paramètres clés et logique

- `src/utils.go` gère les tailles de grille et la détection de victoire.
- `src/game.go` maintient l’état de la partie: grille, tour courant, gagnant, fin de partie, égalité, etc.
  - `GravityEnabled` (bool): vrai si le mode gravité est activé.
  - `GravityUp` (bool): vrai quand la gravité est inversée (les pions « montent »).
  - `Moves` (int): nombre de coups posés; le basculement se produit toutes les 5 poses quand `GravityEnabled` est vrai.
- `pages/play.html` et `templates/*.html` structurent l’UI; `static/play.css` applique l’indicateur.

## ❓ Dépannage

- Rien n’écoute sur 8080 ou port occupé: changez le port dans `src/serveur.go` (ListenAndServe).
- Styles manquants: assurez-vous d’accéder au serveur depuis la racine et que `/static/` est bien servi.
- Erreur de template: vérifiez la présence de `templates/board.html` et `templates/result.html` et que `play.html` les inclut.
- Échec de compilation: utilisez une version récente de Go (>= 1.22) et exécutez la commande depuis la racine du projet.
- Si la grille ne s'affiche pas correctement dans votre navigateur essayer de reduire le zoom et de vider le cache.


