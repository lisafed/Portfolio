// Sélection du bouton pour basculer entre les thèmes
const themeToggleButton = document.getElementById('theme-toggle');

// Stocker les préférences de l'utilisateur dans le localStorage
let currentTheme = localStorage.getItem('theme');

// Si une préférence existe déjà, appliquez-la
if (currentTheme) {
    document.body.classList.add(currentTheme);
    themeToggleButton.textContent = currentTheme === 'dark-mode' ? 'Mode Clair' : 'Mode Sombre';
} else {
    document.body.classList.add('light-mode');
    themeToggleButton.textContent = 'Mode Sombre';
}

// Écouter le clic sur le bouton de basculement de thème
themeToggleButton.addEventListener('click', function () {
    // Vérifier le mode actuel
    if (document.body.classList.contains('light-mode')) {
        // Passer en mode sombre
        document.body.classList.replace('light-mode', 'dark-mode');
        themeToggleButton.textContent = 'Mode Clair';
        localStorage.setItem('theme', 'dark-mode'); // Enregistrer la préférence dans localStorage
    } else {
        // Passer en mode clair
        document.body.classList.replace('dark-mode', 'light-mode');
        themeToggleButton.textContent = 'Mode Sombre';
        localStorage.setItem('theme', 'light-mode'); // Enregistrer la préférence dans localStorage
    }
});
