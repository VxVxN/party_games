package neverhaveiever

import (
	"net/http"

	"github.com/VxVxN/party_games/pkg/httptools"
)

var descriptionText = "\nИгра «Я никогда не» — это захватывающий и увлекательный способ провести время на любой вечеринке, позволяющий не только развлекаться, но и глубже узнать новых людей в компании, а также открыть интересные и порой неожиданные детали о старых друзьях. Существует две версии этой игры: алкогольная и безалкогольная. Каждая из них предлагает уникальный опыт взаимодействия и веселья. \n\nДавайте более подробно рассмотрим правила безалкогольной версии. Перед началом игры участники должны продемонстрировать свои десять пальцев. Эти пальцы представляют собой десять очков, и их количество служит индикатором оставшихся «жизней» в игре. Удобно устроившись — можно положить руки на колени, стол или пол — игроки готовы к началу.\n\nВедущий, который держит телефон, произносит фразу: «Я никогда не…». Если у кого-то из участников имеется опыт совершения того или иного действия, о котором говорит ведущий, этот игрок аккуратно загибает один палец. Это правило распространяется на всех, включая самого ведущего, что делает игру ещё более увлекательной и живой. \n\nЗатем ведущий зачитывает следующее предложение, и процесс повторяется. Игра продолжается до тех пор, пока не останется лишь один человек с прямыми пальцами. Этот игрок становится победителем.\n\nЧто касается алкогольной версии, здесь правила могут быть немного изменены. Когда участник загибает пальц ему нужно выпить определённое колличество алкогольного напитка. Можно также делать это без загибания пальцев, просто наслаждаться игрой и напитками и весело проводить время."

// DescriptionHandler Return description of "I have never" game
//
//	@Summary		Description of game
//	@Description	get description
//	@Tags			neverhaveiever
//	@Produce		json
//	@Success		200	{object}	httptools.JsonSuccessResponse{result=string}	"desc"
//	@Router			/neverhaveiever/description [get]
func (controller *Controller) DescriptionHandler(w http.ResponseWriter, r *http.Request) {
	httptools.SuccessResponse(w, descriptionText)
}
