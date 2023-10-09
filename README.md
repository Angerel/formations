# Goal

This repository aims at exposing the diverse projects that I have been creating to improve my knowledge.\
It also allows one to list all the languages I am proficient in.

Each folder will be centered around one language or one framework, whether it is frontend, backend or for another use.
To serve this purpose, more information will be delivered in each folder, with another README

Furthermore, each and every frontend will display a todolist application, along with a live chat.
Paired with the frontend, any backend will expose a REST API providing the todolist endpoints, and a Websocket to give the live messages

# GitHub organization

Each language/framework will have a main branch dedicated to it, through the naming `origin/<folder>/main`.

For example, Go folder will have its main branch named `origin/go/main`.\
Any branch adding feature would be then named like this : `origin/go/feature/<featureName>`, and any branch correcting a bug would be named `origin/go/fix/<bugName>`.

To validate a change, the "feature" and "fix" branches must be merged into the sub-main branches first, and only the sub-main branches can be merged into the main branch.
