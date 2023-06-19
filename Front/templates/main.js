$(document).ready(function() {
    // Manipulez le formulaire d'enregistrement lors de la soumission
    $("#register-form").submit(function(event) {
        event.preventDefault(); // Empêche le rechargement de la page

        // Récupérez les valeurs des champs de formulaire
        var email = $("#email").val();
        var password = $("#password").val();
        var confirmPassword = $("#confirm-password").val();

        // Vérifiez si les mots de passe correspondent
        if (password !== confirmPassword) {
            alert("Les mots de passe ne correspondent pas");
            return;
        }

        // Créez un objet contenant les données d'enregistrement
        var userData = {
            email: email,
            password: password
        };

        // Effectuez une requête AJAX pour envoyer les données au serveur
        $.ajax({
            url: "http://localhost:8000/api/register", // Remplacez l'URL par celle de votre API
            type: "POST",
            data: JSON.stringify(userData),
            contentType: "application/json",
            success: function(response) {
                // Traitement de la réponse du serveur en cas de succès
                console.log(response);
                alert("Enregistrement réussi !");
            },
            error: function(error) {
                // Traitement de l'erreur en cas d'échec de la requête
                console.log(error);
                alert("Erreur lors de l'enregistrement.");
            }
        });
    });
});
