# TP - Algorithmique de tri en GO<br>
--------------------------------------<br>

lien github : https://github.com/Juuuules83/tp-tris.git<br>

--------------------------------------<br>

## EX01 - Le tri à bulles :<br>

### Question : mesurez votre tri sur SortedScores(10 000), puis sur ReversedScores(10 000). Expliquez l'écart.<br>
### Que se passerait-il si votre tri ne s'arrêtait pas après un passage sans échange ?<br>

Pour SortedScores(10 000), le tableau est déjà trié. Le tri fait donc un seul passage sans échange et s'arrête rapidement.<br>

Pour ReversedScores(10 000), le tableau est dans l'ordre inverse. Le tri doit donc effectuer beaucoup plus d'échanges et de passages, ce qui explique que le temps d'exécution soit beaucoup plus élevé.<br>

La complexité du tri à bulles est de O(n²) dans le pire cas. Pour un tableau déjà trié, elle est de O(n) grâce à l'arrêt après un passage sans échange.<br>

Si le tri ne s'arrêtait pas après un passage sans échange, il continuerait à parcourir le tableau alors qu'il est déjà trié. Il ferait donc des comparaisons inutiles et prendrait plus de temps.<br>

**Commande pour lancer le bench :** _go test -bench=BubbleSort -benchmem -run='^$'_ <br>

