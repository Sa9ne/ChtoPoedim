const API_AUTH = "http://localhost:8081";

// Переместили функции в глобальную область
function toggleAuthForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (authForm && registerForm) {
        if (authForm.classList.contains("show") || registerForm.classList.contains("show")) {
            authForm.classList.remove("show");
            registerForm.classList.remove("show");
        } else {
            authForm.classList.add("show");
        }
    }
}

document.addEventListener("DOMContentLoaded", function () {
    const openAuthBtn = document.querySelector(".avatar");

    if (openAuthBtn) {
        openAuthBtn.addEventListener("click", function (event) {
            event.stopPropagation();
            toggleAuthForm();
        });
    }
});

function toggleToRegisterForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (authForm) authForm.classList.remove("show");
    if (registerForm) registerForm.classList.add("show");
}

document.addEventListener("DOMContentLoaded", function () {
    const registerLink = document.getElementById("register-link");

    if (registerLink) {
        registerLink.addEventListener("click", function (event) {
            event.preventDefault();
            toggleToRegisterForm();
        });
    }
});

function toggleToLoginForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (registerForm) registerForm.classList.remove("show");
    if (authForm) authForm.classList.add("show");
}

document.addEventListener("DOMContentLoaded", function () {
    const backToLogin = document.getElementById("back-to-login");

    if (backToLogin) {
        backToLogin.addEventListener("click", function (event) {
            event.preventDefault();
            toggleToLoginForm();
        });
    }
});

// Функция для переключения между формами регистрации и входа
function toggleRegisterForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (authForm) {
        authForm.classList.remove("show");
    }
    if (registerForm) {
        registerForm.classList.add("show");
    }
}


document.addEventListener("DOMContentLoaded", function () {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");
    const profileForm = document.getElementById("profile-form");
    const openAuthBtn = document.querySelector(".avatar"); // Кнопка для входа
    const closeBtns = document.querySelectorAll(".close-btn"); // Кнопки закрытия
    const registerLink = document.getElementById("register-link");
    const backToLogin = document.getElementById("back-to-login");
    const registerButton = document.getElementById("register-button");
    const authButton = document.getElementById("auth-button");
    const quitButton = document.getElementById("quit-button");

    // Открытие формы входа по клику на аватар
    if (openAuthBtn) {
        openAuthBtn.addEventListener("click", function (event) {
            event.stopPropagation();
            toggleAuthForm();
        });
    }

    // Закрытие форм при клике вне их области
    document.addEventListener("click", function (event) {
        if (
            authForm && registerForm && openAuthBtn &&
            !authForm.contains(event.target) &&
            !registerForm.contains(event.target) &&
            !openAuthBtn.contains(event.target)
        ) {
            authForm.classList.remove("show");
            registerForm.classList.remove("show");
        }
    });

    // Остановка всплытия события внутри форм
    if (authForm) authForm.addEventListener("click", (event) => event.stopPropagation());
    if (registerForm) registerForm.addEventListener("click", (event) => event.stopPropagation());

    // Закрытие форм по кнопкам
    closeBtns.forEach(btn => {
        btn.addEventListener("click", function () {
            if (authForm) authForm.classList.remove("show");
            if (registerForm) registerForm.classList.remove("show");
        });
    });

    // Обработчики переключения форм
    if (registerLink) registerLink.addEventListener("click", (e) => {
        e.preventDefault();
        toggleToRegisterForm();
    });

    if (backToLogin) backToLogin.addEventListener("click", (e) => {
        e.preventDefault();
        toggleToLoginForm();
    });

    // Функция регистрации пользователя
    async function registerUser(event) {
        event.preventDefault();

        const username = document.querySelector("#register-form input[placeholder='Username']").value;
        const email = document.querySelector("#register-form input[placeholder='Email']").value;
        const password = document.querySelector("#register-form input[placeholder='Password']").value;

        const response = await fetch(`${API_AUTH}/api/register`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, email, password })
        });

        const data = await response.json();

        if (response.ok) {
            localStorage.setItem("token", data.token);
            showProfile(username);
        } else {
            alert("Ошибка регистрации: " + (data.error || "неизвестная ошибка"));
        }
    }

    if (registerButton) {
        registerButton.addEventListener("click", registerUser);
    }

    // Функция входа пользователя
    async function loginUser(event) {
        event.preventDefault();

        const email = document.querySelector("#auth-form input[placeholder='Email']").value;
        const password = document.querySelector("#auth-form input[placeholder='Password']").value;

        const response = await fetch(`${API_AUTH}/api/login`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password })
        });

        const data = await response.json();

        if (response.ok) {
            localStorage.setItem("token", data.token);
            showProfile(data.username);
        } else {
            alert("Ошибка входа: " + (data.error || "неизвестная ошибка"));
        }
    }

    if (authButton) {
        authButton.addEventListener("click", loginUser);
    }

    // Проверка авторизации
    async function checkAuth() {
        const token = localStorage.getItem("token");
        if (!token) return;

        const response = await fetch(`${API_AUTH}/api/profile`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        });

        if (response.ok) {
            const data = await response.json();
            showProfile(data.username);
        } else {
            localStorage.removeItem("token");
        }
    }

    checkAuth();

    // Функция отображения профиля
    function showProfile(username) {
        const authForm = document.getElementById("auth-form");
        const registerForm = document.getElementById("register-form");
        const profile = document.getElementById("profile-form");
        const nickname = document.getElementById("nickname");

        if (authForm) authForm.classList.add("hidden");
        if (registerForm) registerForm.classList.add("hidden");
        if (profile) profile.classList.remove("hidden");
        if (nickname) nickname.textContent = username;
    }

    // Функция выхода из аккаунта
    function quitAccount() {
        localStorage.removeItem("token");
        const profile = document.getElementById("profile-form");
        const authForm = document.getElementById("auth-form");

        if (profile) profile.classList.add("hidden");
        if (authForm) authForm.classList.remove("hidden");
    }

    if (quitButton) {
        quitButton.addEventListener("click", quitAccount);
    }
});
