package structures
 
import (
	"encoding/json"
	"os"
)
 
const FichierSauvegarde = "sauvegarde.json"
 
// BONUS : sauvegarde le personnage dans un fichier JSON
func Sauvegarder(p Personne) error {
	donnees, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FichierSauvegarde, donnees, 0644)
}
 
// BONUS : charge le personnage depuis le fichier de sauvegarde.
// Retourne (Personne{}, false) si aucune sauvegarde n'existe ou si elle est illisible.
func Charger() (Personne, bool) {
	donnees, err := os.ReadFile(FichierSauvegarde)
	if err != nil {
		return Personne{}, false
	}
 
	var p Personne
	if err := json.Unmarshal(donnees, &p); err != nil {
		return Personne{}, false
	}
 
	return p, true
}
 
// BONUS : vérifie si un fichier de sauvegarde existe déjà
func SauvegardeExiste() bool {
	_, err := os.Stat(FichierSauvegarde)
	return err == nil
}
 