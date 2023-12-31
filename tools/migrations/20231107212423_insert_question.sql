-- +goose Up
-- +goose StatementBegin
INSERT INTO question (question_id, question_text, answer_1, answer_2, answer_3, answer_4, next_question)
VALUES
    (0, 'Привет, я могу помочь тебе с выбором парфюмерии', 'Привет! Давай начнем', '', '', '', 1),
    (1, 'Какой у вас бюджет?', '', '', '', '', 2),
    (2, 'Для чего вам духи?', 'Для себя', 'Подарок', '', '', 3),
    (3, 'Какую ассоциацию или эмоцию вы хотите вызвать у человека, которому предназначается подарок?', 'Счастье', 'Умиротворение', 'Вдохновение', 'Роскошь', 4),
    (4, 'Какого возраста адресат подарка? Духи для молодежи или для более зрелых людей?', 'Для молодежи', 'Взрослый человек', '', '', 5),
    (5, 'Кому вы дарите духи?', 'Девушке', 'Парню', 'Другу', 'Родственнику', 21),
    (6, 'Какой ваш возраст?', '', '', '', '', 7),
    (7, 'Какой ваш пол?', 'Мужчина', 'Женщина', '', '', 8),
    (8, 'Какой у вас тип кожи?', 'Жирная', 'Сухая', 'Комбинированная', '', 9),
    (9, 'Когда вы будете использовать духи?', 'Для повседневного пользования', 'На работе', 'Во время отдыха', '', 10),
    (10, 'В какое время года вы будете использовать духи?', 'Лето', 'Зима', 'Весна/Осень', 'В любое', 11),
    (11, 'Какой тип ароматов вам нравится больше всего?', 'Сладкий', 'Душистый', 'Крепкий', '', 12),
    (12, 'Предпочитаете ли вы легкие и свежие ароматы или насыщенные и тяжелые?', 'Легкие и свежие', 'Насыщенные и тяжелые', '', '', 13),
    (13, 'Что важнее для вас: долгое время удержания аромата или интенсивность запаха?', 'Удержание аромата', 'Интенсивность запаха', '', '', 14),
    (14, 'Какую атмосферу вы хотите создавать с помощью аромата: романтическую, элегантную, энергичную или провокационную?', 'Романтическую', 'Элегантную', 'Энергичную', 'Провокационную', 15),
    (15, 'Предпочитаете ли вы классические ароматы или что-то более современное и эксклюзивное?', 'Классические ароматы', 'Современные и экзлюзивное', '', '', 16),
    (16, 'Какой тип ароматов вам нравится больше всего: цветочные, фруктовые, древесные или ориентальные?', 'Цветочные', 'Древесные', 'Фруктовые', 'Ориентальные', 17),
    (17, 'Предпочитаете ли вы потрясающий и привлекательный аромат или что-то более непритязательное и ненавязчивое?', 'Привлекательный', 'Ненавязчивый', '', '', 18),
    (18, 'Что вы думаете о сезонных ограничениях ароматов?', 'Они дополняют атмосферу каждого сезона', 'Они не имеют большого значения', 'Они ограничивают свободу выбора', '', 19),
    (19, 'Какую стойкость и проекцию аромата вы ожидаете от духов?', 'Сильная стойкость', 'Чтобы держались недолго', '', '', 20),
    (20, 'Какой объем духов вас интересует?', '100 мл', '50 мл', '10 мл', '', 21),
    (21, 'Тест завершен! Вы можете посмотреть результаты', 'Посмотреть результат', '', '', '', -1)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM question;
-- +goose StatementEnd
