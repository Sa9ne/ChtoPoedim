document.addEventListener("DOMContentLoaded", function () {
    const langModal = document.getElementById("lang-modal");
    const langButton = document.getElementById("selected-lang");

    if (!langModal || !langButton) {
        console.error("Ошибка: не найден элемент с id lang-modal или selected-lang.");
        return;
    }

    // Функция для открытия/закрытия модального окна
    function toggleLangModal() {
        langModal.style.display = (langModal.style.display === "block") ? "none" : "block";
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
                // Сохранение языка в cookies
                document.cookie = `lang=${lang}; path=/; max-age=3600`;

                // Обновление текста на странице
                document.querySelectorAll("[data-translate]").forEach(el => {
                    const key = el.getAttribute("data-translate");

                    if (translations[key]) {
                        if (el.tagName === "INPUT") {
                            el.placeholder = translations[key]; // Для input обновляем placeholder
                        } else {
                            el.innerHTML = translations[key]; // Меняем текст без удаления вложенных тегов
                        }
                    }
                });

                // Обновляем текст кнопки выбора языка
                updateLanguageButton();

                // Закрываем модальное окно после смены языка
                langModal.style.display = "none";
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

    // Обновление текста кнопки выбора языка
    function updateLanguageButton() {
        const lang = getCookie("lang") || "en";
        langButton.textContent = lang === "ru" ? "Рус" : "Eng";
    }

    // Закрытие модального окна при клике вне его
    document.addEventListener("click", function (event) {
        if (
            langModal.style.display === "block" &&
            event.target !== langModal &&
            event.target !== langButton &&
            !langModal.contains(event.target)
        ) {
            langModal.style.display = "none";
        }
    });

    // Инициализация текста кнопки при загрузке
    updateLanguageButton();
    langModal.style.display = "none";

    // Добавляем функции в глобальную область видимости, чтобы они работали с HTML
    window.toggleLangModal = toggleLangModal;
    window.changeLanguage = changeLanguage;
});
