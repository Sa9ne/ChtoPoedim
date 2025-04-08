let CurrentLang = "en";
let translations = {};
// Функция открытия окна языков
function OpenLangModal() {
    CloseModal(OpenLang);
    OpenLang.classList.toggle("hidden")
}
// Обновление текста кнопки
function updateLanguageButton() {
    const langButton = document.getElementById("lang-button");
    if (CurrentLang === "en") {
        langButton.innerText = "Eng";  // Текст для английского языка
    } else if (CurrentLang === "ru") {
        langButton.innerText = "Рус";  // Текст для русского языка
    }
}
// Запрос к серверу для языка 
async function LoadTranslations(lang) {
    try {
        const response = await fetch(`/frontend/lang/${lang}.json`);
        translations = await response.json();
        applyTranslations();
    } catch(err) {
        console.error("Translations doesn't found", err)
    }
}
// Функция, которая применяет переводы ко всем элементам с атрибутом data-translate
function applyTranslations() {
    document.querySelectorAll('[data-translate]').forEach(el => {
        // Получаем значение атрибута (перевод)
        const key = el.getAttribute('data-translate');
        // Получаем переведённый текст из объекта translations по ключу
        const text = translations[key];
        // Если элемент не просто текст
        if (el.tagName === 'INPUT' || el.tagName === 'TEXTAREA') {
            el.placeholder = text || "";
        } else {
            el.innerHTML = text || "";
        }
    });
}
// Функция выбора языка
async function changeLanguage(lang) {
    try {
        OpenLangModal();
        await LoadTranslations(lang);
        applyTranslations();

        localStorage.setItem('language', lang);
        CurrentLang = lang;
        updateLanguageButton();
    } catch (err) {
        console.error("Error while changing language:", err);
    }
}
// При загрузке страницы проверяем, был ли выбран язык, и загружаем соответствующие переводы
window.addEventListener('load', () => {
    const savedLang = localStorage.getItem('language') || 'en';
    if (savedLang !== CurrentLang) {
        changeLanguage(savedLang);
        updateLanguageButton(); 
        CloseModal();
    }
});