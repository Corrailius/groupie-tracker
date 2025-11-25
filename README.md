# groupie tracker

---

### ✅ 1. Un commit = une seule modification logique

Évite les commits fourre-tout. Chaque commit doit avoir un but clair.

### ✅ 2. Messages courts et précis

Format recommandé :
**type: description courte**
Exemples :

* `feat: ajouter la page de login`
* `fix: corriger l’erreur de connexion`
* `docs: mettre à jour le README`

### ✅ 3. Utiliser l’impératif

Comme si tu donnais un ordre :
✔️ "add login page"
❌ "added login page"
❌ "adding login page"

### ✅ 4. Ajouter un message détaillé si nécessaire

Si la modification est complexe, ajoute un paragraphe expliquant le “pourquoi”.

### ✅ 5. Toujours vérifier avant de commit

* Pas de fichiers inutiles
* Pas de mots de passe / secrets
* Code formaté

### ✅ 6. Commits fréquents mais propres

Ne garde pas tout en local pendant 3 jours. Commit et pousse régulièrement.

### Exemple de bon commit :

```
fix: corrige le bug d'affichage du menu sur mobile

Le problème venait d’un mauvais breakpoint CSS. Ajustement de la media query.
```
