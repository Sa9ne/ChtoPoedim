// Функция для открытия/закрытия модального окна
function toggleLangModal() {
    const modal = document.getElementById("lang-modal");

    if (modal.style.display === "none" || modal.style.display === "") {
        modal.style.display = "block";
    } else {
        modal.style.display = "none";
    }
}

// Функция смены языка
function changeLanguage(lang) {
    fetch(`/get-translations?lang=${lang}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Ошибка сервера: ${response.status}`);
            }
            return response.json();
        })
        .then(translations => {
            // Сохраняем язык в cookies
            document.cookie = `lang=${lang}; path=/; max-age=3600`;

            // Обновляем текст на странице
            document.querySelectorAll("[data-translate]").forEach(el => {
                const key = el.getAttribute("data-translate");

                if (translations[key]) {
                    if (el.tagName === "INPUT") {
                        el.placeholder = translations[key]; // Для input обновляем placeholder
                    } else if (el.tagName === "A") {
                        el.innerHTML = translations[key]; // Для ссылок сохраняем теги <a>
                    } else {
                        el.innerHTML = translations[key]; // Меняем текст без удаления вложенных тегов
                    }
                }
            });

            // Обновляем кнопку с выбранным языком
            updateLanguageButton();
        })
        .catch(error => console.error("Ошибка при смене языка:", error));
}

// Получение cookie по имени
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
    return null;
}

// Обновление текста кнопки на текущий язык
function updateLanguageButton() {
    const lang = getCookie("lang") || "en";
    const button = document.getElementById("selected-lang");

    if (button) {
        button.textContent = lang === "ru" ? "Рус" : "Eng";
    }
}

// Закрытие модального окна при клике вне его
document.addEventListener("click", function (event) {
    const modal = document.getElementById("lang-modal");
    const button = document.getElementById("selected-lang");

    if (modal.style.display === "block" && event.target !== modal && event.target !== button && !modal.contains(event.target)) {
        modal.style.display = "none";
    }
});

// Устанавливаем правильный язык в кнопке при загрузке страницы
document.addEventListener("DOMContentLoaded", () => {
    updateLanguageButton();
    document.getElementById("lang-modal").style.display = "none";
});

document.getElementById("selected-lang").addEventListener("click", function (event) {
    toggleLangModal(); // При клике на кнопку открываем или закрываем модальное окно
});