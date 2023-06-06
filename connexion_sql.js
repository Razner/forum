/*
// Lecture du fichier SQL 
const fs = require('fs');
const sqlScript = fs.readFileSync('forum.sql', 'utf8');

// Execution des requêtes SQL
db.exec (sqlScript, function (err) {
    if (err) {
        console.error(err.message);
    } else {
        console.log('Création des Tables réussies.');
    }

    // Fermeture de la connexion à la base de données
    db.close();
});*/