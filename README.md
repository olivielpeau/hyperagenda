# Hyper Agenda

## Kezako?

**Hyper Agenda**, parce qu'importer son Hyper Planning dans *Google Agenda*, c'est la base.

### Comment ça marche

* Récupérez l'url ical dans hyperplanning 

  ![Screenshot](http://i.imgur.com/oQdfm7q.png)

* Remplacez *"agenda.ecp.fr"* par *"hyperagenda.herokuapp.com"*
* Utilisez cette url pour l'importation dans google agenda

![Screenshot](http://i.imgur.com/Qm30jHz.png)

#### FAQ

**Ça marche pas…**

Si votre url hyperplanning ne commence pas par *"/ebauche/Telechargements/ical"* ça ne vas pas marcher, question de sécurité. Contactez moi si ça devrait.

**Et la NSA dans tout ça?**

Je ne fais rien de votre url (qui contiens un token Hyper Planning), allez voir le code source si vous voulez. L'appli est hébergée par Heroku en Europe. Les logs http sont supprimés régulièrement.


---

Made w/ &#x2665; by [Adrien Kohlbecker](mailto:adrien.kohlbecker@gmail.com?subject=[HyperAgenda]%20). [Source Code](https://github.com/adrienkohlbecker/hyperagenda)