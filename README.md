Write a dice game script that accepts N number of players and M number of players as input
dice, with the following rules:
1. At the start of the game, each player gets a dice of M units.
2. All players will roll their respective dice at the same time
3. Each player will check the results of their dice rolls and evaluate
as follows:<br>
    a. Dice number 6 will be removed from the game and added as points for the player.<br>
    b. Dice number 1 will be awarded to the player sitting next to him.<br>
       For example, the first player will give his dice the number 1 to the second player.<br>
    c. Dice numbers 2,3,4 and 5 will still be played by the player.<br>
4. After evaluation, the player who still has the dice will repeat the 2nd step until only 1 player remains.
   a Players who have no more dice are considered to have finished playing.
5. The player who has the most points wins.