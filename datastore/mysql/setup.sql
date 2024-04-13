insert into payment_details (id, description)
values (1, 'buy_assistant_coach'),
       (2, 'buy_doctor'),
       (3, 'buy_fitness_coach'),
       (4, 'buy_talent_finder'),
       (5, 'buy_trainer'),
       (6, 'buy_player'),
       (7, 'salaries'),
       (8, 'sell_assistant_coach'),
       (9, 'sell_doctor'),
       (10, 'sell_fitness_coach'),
       (11, 'sell_talent_finder'),
       (12, 'sell_trainer'),
       (13, 'sell_player'),
       (14, 'player_salaries');


insert into doctor_levels (id, healing, max_team_players, salary)
values (1),
       (2),
       (3),
       (4),
       (5),
       (6);


insert into talent_finder_levels (id, observe_players, observation_time,
                                  encryption_transfer, weekly_transfer_capacity, salary, max_foreigner_players)
values (1),
       (2),
       (3),
       (4),
       (5),
       (6);



insert into trainer_levels (id, salary)
VALUES (1),
       (2),
       (3),
       (4),
       (5),
       (6);

insert into assistant_coach_levels (id, salary)
VALUES (1),
       (2),
       (3),
       (4),
       (5),
       (6);

insert into fitness_coach_levels (id, salary, power_increase_per_day)
VALUES (1),
       (2),
       (3),
       (4),
       (5),
       (6);