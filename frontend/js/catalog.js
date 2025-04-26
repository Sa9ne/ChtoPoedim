// Функция открытия каталога
function OpenCatalog() {
  CloseModal(Catalog);
  Catalog.classList.toggle("hidden")
}
// Функция открытия окна блюда дня
function OpenDayDish() {
  CloseModal(DishDay);
  DishDay.classList.toggle("hidden")
  DayDish();
}
// Функция открытия окна выбора блюд
function OpenSmartSelect() {  
  if (!InfoSite.classList.contains("hidden")){
    InfoSite.classList.toggle("hidden")
  }

  if (!FoodCatalog.classList.contains("hidden")) {
    FoodCatalog.classList.add("hidden");
  }

  SmartSelect.classList.toggle("hidden")
}
// Функция открытия каталога еды
function OpenFoodCatalog() {
  if (!InfoSite.classList.contains("hidden")){
    InfoSite.classList.toggle("hidden")
  }

  if (!SmartSelect.classList.contains("hidden")) {
    SmartSelect.classList.add("hidden");
  }

  FoodCatalog.classList.toggle("hidden")
}
// Функция загрузки перевода на русский
let translations_dish = {};

async function LoadTranslation(lang) {
  try {
    const response = await fetch(`/frontend/lang/dish/${lang}.json`);
    if (!response.ok) throw new Error("Ошибка загрузки перевода");
    translations_dish = await response.json();
  } catch (error) {
    console.error("Loading translation problem")
  }
}
// Функция перевода на русский
function translate(text) {
  return translations_dish[text] || text;
}
// Функция блюда дня
async function DayDish() {
  try {
    const response = await fetch("http://localhost:8082/DishDay");
    if (!response.ok) throw new Error("Error found dish")

    const lang = localStorage.getItem("language")
    await LoadTranslation(lang);

		const dish = await response.json();

		// Заполняем содержимое
		DishName.textContent = translate(dish.name) || "No name";
		DishDescription.textContent = translate(dish.description) || "Null description";
		DishTime.textContent = dish.time || "-";
  }  catch (error) {
		console.error("Ошибка:", error);
	}
}