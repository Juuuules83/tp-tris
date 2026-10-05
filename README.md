# TP - Algorithmique de tri en GO<br>
-------------------------------------<br>

lien github : https://github.com/Juuuules83/tp-tris.git<br>

-------------------------------------<br>

## EX01 - Le tri à bulles :<br>
### Question : mesurez votre tri sur SortedScores(10 000), puis sur ReversedScores(10 000). Expliquez l'écart.<br>
### Que se passerait-il si votre tri ne s'arrêtait pas après un passage sans échange ?

voici le résultat du benchmark <br>
-> ![alt text](image.png)

On constate que le temps d'exécution augmente très fortement quand le nombre de scores augmente.<br>

Pour 1 000 éléments, le tri prend environ 0,37 ms, alors que pour 10 000 éléments il prend environ 35,6 ms et pour 100 000 éléments environ 9,85 secondes.<br>

La complexité du tri à bulles est de O(n²) dans le pire cas... quand le nombre d'éléments augmente, le nombre de comparaisons et d'échanges augmente donc eux aussi.<br>

Pour un tableau déjà trié, le tri peut s'arrêter après un passage sans échange. Si le tri ne s'arrête pas, il continuerait à parcourir le tableau alors qu'il est déjà trié et ferait donc des comparaisons qui ne servirait pas.<br>

**Commande pour lancer le bench :** _go test -bench=BubbleSort -benchmem -run='^$'_  <br>

--------------------------------------<br>

## EX02 - Le tri par sélection :<br>

voici le résultat du benchmark V1<br>
->  


**Commande pour lancer le bench :** _ // _  <br>

--------------------------------------<br>
