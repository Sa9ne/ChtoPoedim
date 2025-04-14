// Функция открытия каталога
function OpenCatalog() {
  CloseModal(Catalog);
  Catalog.classList.toggle("hidden")
}
// Функция открытия окна блюда дня
function OpenDayDish() {
  CloseModal(DishDay);
  DishDay.classList.toggle("hidden")
}