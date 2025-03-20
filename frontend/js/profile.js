// Функция для открытия и закрытия модального окна
function toggleProfileForm() {
    const form = document.getElementById("profile-form");
    form.classList.toggle("hidden");
}

// Функция для выхода (Quit) пользователя
function quitAccount() {
    localStorage.removeItem("auth_token");

    window.location.href = "/login";
}

// Функция для получения данных пользователя из токена или API
function fetchUserProfile() {
    const token = localStorage.getItem("auth_token");

    if (token) {
        fetch("http://localhost:8081/profile", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
        .then(response => response.json())
        .then(data => {
            if (data.nickname && data.favorite_food) {
                document.getElementById("nickname").value = data.nickname;
                document.getElementById("favorite-food").value = data.favorite_food;

                toggleProfileForm();
            }
        })
        .catch(error => console.error("Ошибка получения данных профиля", error));
    }
}

fetchUserProfile();
