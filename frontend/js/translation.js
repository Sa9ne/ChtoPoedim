// Функция для открытия/закрытия модального окна
function toggleLangModal() {
    const modal = document.getElementById("lang-modal");
    modal.style.display = modal.style.display === "block" ? "none" : "block";
}

// Получение cookie по имени
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

// Обновление текста кнопки на текущий язык
function updateLanguageButton() {
    const lang = getCookie("lang") || "en";
    const button = document.getElementById("selected-lang");

    if (lang === "en") {
        button.innerHTML = "Eng";
    } else if (lang === "ru") {
        button.innerHTML = "Рус";
    }
}

// Функция смены языка
function changeLanguage(lang) {
    fetch(`/set-languages?lang=${lang}`)
        .then(() => {
            document.cookie = `lang=${lang}; path=/; max-age=3600`;

            // Обновляем кнопку с выбранным языком
            updateLanguageButton();

            // Закрываем модальное окно
            document.getElementById("lang-modal").style.display = "none";

            // Перезагрузка страницы для применения изменений
            location.reload();
        })
        .catch(error => console.error("Ошибка при смене языка:", error));
}

// Закрытие окна при клике вне его
document.addEventListener("click", function (event) {
    const modal = document.getElementById("lang-modal");
    const button = document.getElementById("selected-lang");
    if (event.target !== modal && event.target !== button) {
        modal.style.display = "none";
    }
});

// Устанавливаем правильный язык в кнопке при загрузке страницы
document.addEventListener("DOMContentLoaded", updateLanguageButton);
