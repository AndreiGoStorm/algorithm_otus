package lz77

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	tempPath      = "/tmp/"
	tempExtension = "*.txt"
)

func TestWords(t *testing.T) {
	tests := []struct {
		text string
	}{
		{text: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{text: "ALICE IN WONDERLAND"},
		{text: "ABRAKADABRA"},
		{text: "ABABABABABAC"},
		{text: "abcabcabcabcabcabc"},
		{text: "КОСИЛКОСОЙКОСОЙ"},
		{text: "abcdefghijklmnopqrstuvwxyz"},
		{text: "To be, or not to be, that is the question."},
		{text: "The rain in Spain falls mainly in the plain"},
		{text: "<div><span>HTML</span><span>World</span></div>"},
		{text: "Searching Auto Algorithm, build frequency table"},
		{text: "About half-past ten the cracked bell of the small church began to ring,\nand presently the people began to gather for the morning sermon. The\nSunday-school children distributed themselves about the house and\noccupied pews with their parents, so as to be under supervision. Aunt\nPolly came, and Tom and Sid and Mary sat with her—Tom being placed next\nthe aisle, in order that he might be as far away from the open window\nand the seductive outside summer scenes as possible. The crowd filed up\nthe aisles: the aged and needy postmaster, who had seen better days;\nthe mayor and his wife—for they had a mayor there, among other\nunnecessaries; the justice of the peace; the widow Douglas, fair,\nsmart, and forty, a generous, good-hearted soul and well-to-do, her hill\nmansion the only palace in the town, and the most hospitable and much\nthe most lavish in the matter of festivities that St. Petersburg could\nboast; the bent and venerable Major and Mrs. Ward; lawyer Riverson, the\nnew notable from a distance; next the belle of the village, followed by\na troop of lawn-clad and ribbon-decked young heart-breakers; then all\nthe young clerks in town in a body—for they had stood in the vestibule\nsucking their cane-heads, a circling wall of oiled and simpering\nadmirers, till the last girl had run their gantlet; and last of all came\nthe Model Boy, Willie Mufferson, taking as heedful care of his mother as\nif she were cut glass. He always brought his mother to church, and was\nthe pride of all the matrons. The boys all hated him, he was so\ngood. And besides, he had been “thrown up to them” so much. His\nwhite handkerchief was hanging out of his pocket behind, as usual on\nSundays—accidentally. Tom had no handkerchief, and he looked upon boys\nwho had as snobs.\n\nThe congregation being fully assembled, now, the bell rang once more,\nto warn laggards and stragglers, and then a solemn hush fell upon the\nchurch which was only broken by the tittering and whispering of the\nchoir in the gallery. The choir always tittered and whispered all\nthrough service. There was once a church choir that was not ill-bred,\nbut I have forgotten where it was, now. It was a great many years ago,\nand I can scarcely remember anything about it, but I think it was in\nsome foreign country.\n\n  Shall I be car-ri-ed toe the skies, on flow’ry _beds_\n                                                        of ease,\n\n  Whilst others fight to win the prize, and sail thro’ _blood_\n                                                        -y seas?\n\nHe was regarded as a wonderful reader. At church “sociables” he was\nalways called upon to read poetry; and when he was through, the ladies\nwould lift up their hands and let them fall helplessly in their laps,\nand “wall” their eyes, and shake their heads, as much as to say, “Words\ncannot express it; it is too beautiful, TOO beautiful for this mortal\nearth.”\n\nAfter the hymn had been sung, the Rev. Mr. Sprague turned himself into\na bulletin-board, and read off “notices” of meetings and societies and\nthings till it seemed that the list would stretch out to the crack of\ndoom—a queer custom which is still kept up in America, even in cities,\naway here in this age of abundant newspapers. Often, the less there is\nto justify a traditional custom, the harder it is to get rid of it.\n\nAnd now the minister prayed. A good, generous prayer it was, and went\ninto details: it pleaded for the church, and the little children of the\nchurch; for the other churches of the village; for the village itself;\nfor the county; for the State; for the State officers; for the United\nStates; for the churches of the United States; for Congress; for the\nPresident; for the officers of the Government; for poor sailors, tossed\nby stormy seas; for the oppressed millions groaning under the heel of\nEuropean monarchies and Oriental despotisms; for such as have the light\nand the good tidings, and yet have not eyes to see nor ears to hear\nwithal; for the heathen in the far islands of the sea; and closed with\na supplication that the words he was about to speak mi aboutfkissa well-to-do no kissa the words about to speak might find grace\\nand favor, and be as seed sown in fertile ground, yie"},
		{text: "– Nie wyjdzie stamtąd, mówię wam – powiedział pryszczaty, z przekonaniem\nkiwając głową. – Już godzina i ćwierć, jak tam wlazł. Już po nim.\nMieszczanie, stłoczeni wśród ruin, milczeli wpatrzeni w ziejący w rumowisku\nczarny otwór, w zagruzowane wejście do podziemi. Grubas w żółtym kubraku\nprzestąpił z nogi na nogę, chrząknął, zdjął z głowy wymięty biret.\n– Poczekajmy jeszcze – powiedział, ocierając pot z rzadkich brwi.\n– Na co? – prychnął pryszczaty. – Tam, w lochach, siedzi bazyliszek,\nzapomnieliście, wójcie? Kto tam wchodzi, ten już przepadł. Mało to ludzi tam\npoginęło? Na co tedy czekać?\n– Umawialiśmy się przecie – mruknął niepewnie grubas. – Jakże tak?\n– Z żywym się umawialiście, wójcie – rzekł towarzysz pryszczatego, olbrzym\nw skórzanym, rzeźnickim fartuchu. – A nynie on martwy, pewne to jak słońce\nna niebie. Z góry było wiadomo, że na zgubę idzie, jako i inni przed nim.\nPrzecie on nawet bez zwierciadła polazł, z mieczem tylko. A bez zwierciadła\nbazyliszka nie zabić, każdy to wie."},
		{text: "Использование объектно-ориентированной парадигмы широко используется\nво всем мире. Однако, как выясняется, многие люди еще не освоили этот метод\nразработки проектов. Переработанное издание книги станет полезным как для\nчитателей, уже применяющих объектно-ориентированный подход, так и для тех,\nкто не очень хорошо его понимает.\nЛюди, впервые изучающие объектно-ориентированный анализ и \nпроектирование (object-oriented analysis and design (OOAD)), найдут в книге следующую\nинформацию.\n• Концептуальные основы и перспективы развития \nобъектно-ориентированного подхода.\n• Примеры использования методы 00AD на протяжении всего жизненного\nцикла системы.\n• Введение в систему стандартных обозначений, используемых при \nпроектировании систем и программного обеспечения, — Unified Modeling Language\n(UML 2.0)\nДля опытных проектировщиков, использующих метод 00AD, книга будет\nинтересной по следующим причинам\n• Язык UML 2.0 по-прежнему неизвестен многим опытным \nпроектировщикам. Они найдут в книге новую систему обозначений.\n• С учетом критических замечаний, полученных после выхода в свет \nпредыдущих изданий, в новом варианте книги сделан более сильный акцент на\nмоделировании.\n• Прочтя часть Концепции, читатели смогут понять, \"почему вещи таковы,\nкаковы они есть\" в объектно-ориентированном мире. Многие люди до сих\nпор не имеют представления об эволюции концепций \nобъектно-ориентированного проектирования. Впрочем, даже если читателю уже известны\nнекоторые факты на эту тему, он не может отрицать значимости этой \nинформации при первом изучении предмета.\nТретье издание имеет четыре основных отличий от предыдущих.\n1. За прошедшее время язык UML 2.0 получил официальное признание. Эта\nсистема обозначений изложена в главе 5. Для того чтобы облегчить читателю\nизучение этого языка, мы отделили его основные элементы от элементов\nповышенной сложности.\n2. В новом издании рассмотрены новые предметные области (см. часть III).\nВ частности, рассмотрены абстракции разного уровня, начиная системами\nс высокоуровневой архитектурой и заканчивая системами, основанными на\nиспользовании Web-технологии.\n3. После выхода в свет последнего по времени издания язык C++ претерпел\nновые изменения, связанные с концепциями объектно-ориентированного\nпрограммирования. По мнению читателей подчеркивать эту особенность\nязыка C++ уже нет необходимости. Рынок переполнен книгами,\nпосвященными объектно-ориентированному программированию, а также\nучебниками по языкам программирования, основанным на этой парадигме. По этой\nпричине мы решили изъять из книги большинство фрагментов программ.\n4. В заключение, в ответ на пожелания читателей, мы сосредоточили основное\nвнимание на аспектах моделирования в рамках парадигмы OOAD. В\nтретьей части книги, посвященной приложениям, продемонстрированы\nпримеры использования языка UML, причем в каждой главе рассмотрен один\nиз этапов жизненного цикла проекта."},
	}

	for _, tc := range tests {
		t.Run("build frequency table", func(t *testing.T) {
			// Arrange
			file, _ := os.CreateTemp(tempPath, tempExtension)
			defer os.Remove(file.Name())
			file.WriteString(tc.text)

			compress, err := NewLZ77(file.Name())
			require.NoError(t, err)

			// Act
			err = compress.Compress()
			require.NoError(t, err)

			// Assert
			decompress, err := NewLZ77(compress.to)
			require.NoError(t, err)
			err = decompress.Decompress()
			require.NoError(t, err)

			content, err := os.ReadFile(compress.from)
			require.NoError(t, err)
			_, err = os.ReadFile(compress.to)
			require.NoError(t, err)

			decompressedContent, err := os.ReadFile(decompress.to)
			require.NoError(t, err)
			require.Equal(t, string(content), string(decompressedContent))

			require.NoError(t, os.Remove(compress.to))
			require.NoError(t, os.Remove(decompress.to))
		})
	}
}
