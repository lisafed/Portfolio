// Sélection du bouton pour basculer entre les thèmes
const themeToggleButton = document.getElementById('theme-toggle');

// Stocker les préférences de l'utilisateur dans le localStorage
const currentTheme = localStorage.getItem('theme') || 'light-mode';

// Appliquer le thème actuel en fonction des préférences stockées
document.body.className = currentTheme;
themeToggleButton.textContent = currentTheme === 'dark-mode' ? 'Mode Clair' : 'Mode Sombre';

// Écouter le clic sur le bouton de basculement de thème
themeToggleButton.addEventListener('click', function() {
    // Vérifier le mode actuel
    if (document.body.classList.contains('light-mode')) {
        // Passer en mode sombre
        document.body.classList.replace('light-mode', 'dark-mode');
        themeToggleButton.textContent = 'Mode Clair';
        localStorage.setItem('theme', 'dark-mode'); // Enregistrer le thème dans localStorage
    } else {
        // Passer en mode clair
        document.body.classList.replace('dark-mode', 'light-mode');
        themeToggleButton.textContent = 'Mode Sombre';
        localStorage.setItem('theme', 'light-mode'); // Enregistrer le thème dans localStorage
    }
});


