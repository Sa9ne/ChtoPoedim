document.addEventListener("DOMContentLoaded", function () {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");
    const openAuthBtn = document.querySelector(".avatar"); // Открытие окна входа
    const closeBtns = document.querySelectorAll(".close-btn"); // Все кнопки закрытия

    // Функция для показа/скрытия формы входа
    function toggleAuthForm() {
        authForm.classList.toggle("show");
        registerForm.classList.remove("show"); // Закрываем регистрацию, если открыта
    }

    // Функция для показа/скрытия формы регистрации
    function toggleRegisterForm() {
        registerForm.classList.toggle("show");
        authForm.classList.remove("show"); // Закрываем вход, если открыт
    }

    // Переключение из входа в регистрацию
    function toggleToRegisterForm() {
        authForm.classList.remove("show");
        registerForm.classList.add("show");
    }

    // Переключение из регистрации во вход
    function toggleToLoginForm() {
        registerForm.classList.remove("show");
        authForm.classList.add("show");
    }

    // Открытие окна входа при клике на аватар
    if (openAuthBtn) {
        openAuthBtn.addEventListener("click", function (event) {
            event.stopPropagation();
            toggleAuthForm();
        });
    }

    // Закрытие всех окон при клике вне них
    document.addEventListener("click", function (event) {
        if (!authForm.contains(event.target) && !registerForm.contains(event.target) && !openAuthBtn.contains(event.target)) {
            authForm.classList.remove("show");
            registerForm.classList.remove("show");
        }
    });

    // Остановка закрытия при клике внутри окон
    authForm.addEventListener("click", function (event) {
        event.stopPropagation();
    });

    registerForm.addEventListener("click", function (event) {
        event.stopPropagation();
    });

    // Закрытие окна при нажатии на кнопку "×"
    closeBtns.forEach(btn => {
        btn.addEventListener("click", function () {
            authForm.classList.remove("show");
            registerForm.classList.remove("show");
        });
    });

    // Переключение между окнами по ссылкам
    document.querySelector("[data-translate='register_link']").addEventListener("click", function (event) {
        event.preventDefault();
        toggleToRegisterForm();
    });

    document.querySelector("[data-translate='back_to_sign_in']").addEventListener("click", function (event) {
        event.preventDefault();
        toggleToLoginForm();
    });
});
