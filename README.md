# Clone Factory

![N|Solid](https://raw.githubusercontent.com/haagor/client_ES/master/docs/clones_lego.jpg)

Clone Factory est un projet qui a pour but initial de comparer les performances de Python et de Go. Ce projet s'étend peu à peu en intégrant de nouveaux (dans le projet) langages ou outils dans ses comparaisons, comme [Pypy](https://pypy.org/) ou Spark (`todo`).
Cette comparaison de performance se fait avec un cas d'usage d'[Elasticsearch](https://www.elastic.co/fr/). Cela me permet dans le même temps de développer des POC de clients ES, tout en comparant des outils que je rencontre dans mon metier.

J'ai développé deux clients ES, un codé en Python et l'autre en GO. Le client Go n'est pas parallélisé pour être comparé au client Python. La parallélisation du client Go aura un sens dans le cadre de la comparaison avec Spark. 

### Installation

`todo`

### Résultats

![N|Solid](https://raw.githubusercontent.com/haagor/client_ES/master/docs/graphe_compare.png)

Ce premier graphe resulte de 100 éxécutions du client Python, en retirant les 10 temps les plus bas et haut. La démarche est similaire pour l'éxécution avec Pypy et Go.
Ainsi Go affiche de meilleurs performance que Python ou Pypy. Python étant un langage interprété contre Go qui est compilé, la différence de performance s'explique à priori simplement. En revanche le peu de différence entre Python et Pypy est surprenant dans la mesure où Pypy est sensé palier ce défaut du langage interprété, je ménerai l'enquête (`todo`).
J'ai dans un second temps observé les résultats de mes tests et j'ai été plutôt étonné de la dispersion des résultats. Cela a motivé les graphes suivants :
[![N|Solid](https://raw.githubusercontent.com/haagor/client_ES/master/docs/graphe_iter.png)]

La dispersion des temps pour les exécutions du client Go est en effet troublant et mérite là encore enquête de ma part (`todo`).

![N|Solid](https://raw.githubusercontent.com/haagor/client_ES/master/docs/htop_py.png)
![N|Solid](https://raw.githubusercontent.com/haagor/client_ES/master/docs/mac.png)

La répartition de charge sur les coeurs coeurs est aussi surprenante pour moi -> enquête (`todo`).