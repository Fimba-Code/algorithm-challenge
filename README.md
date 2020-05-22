<h1 align="center" style="color:">Desafio de Algoritmos básicos: #1 <h1/>
  
## Introdução 📜

Este desafio é desenhado para programadores iniciantes, e também os mais experientes que desejam acrescentar os seus conhecimentos acerca da resolução de alguns algoritmos em programação.  As soluções não estaram restritas á nenhuma linguagem de programação especifica, os nossos **moderadores** encorajam todos os participantes a usar as línguages de programação de sua perferência. 

## Submissão 🚀

As soluções desenvolvidas devem ser partilhadas no canal [**#code-drop** no discord](https://discord.gg/XDPbSUN), posteriormente, as melhores soluções seram colocadas neste repo e os devs receberão kudos no twitter ou outra reded social de sua perferencia. 

**Dicas importantes**

1. Não te esqueças de formatar o teu código ao antes enviar à communidade, para mais informações visite => [discord](https://support.discord.com/hc/en-us/articles/210298617-Markdown-Text-101-Chat-Formatting-Bold-Italic-Underline).
2. Use funções no teu programa sempre que necessário para a reutilização de códigos e simplificação dos programas. 
3. Resolva o teus exercícios streaming live (ajuda-te a ganhar mais pontos)
4. Desenvolva e partilhe o seu pseudocódigo(algoritimo) junto da tua submissão.


## Challenge 🥋

Os challenges (desafios) começaram do mais simples ao mais arduo, mas poderás resolver-los de forma aleatória. Esse challenge será componsto por 3 problemas sendo que os mesmos ajudaram o participante á desenvolver algum skill especifico.

### Challenge #1 - Palíndromo

Palíndromos é um número/palavra (ou um conjunto de números/palavras) que **lendo de esquerda para a direita ou da direita para esquerda são iguais**. Para este challenge, o teu programa deve receber um input do usuário e definir se este input é ou não um palindromo. 

```sh
Input: SOCORRAM-ME! SUBI NO ÔNIBUS EM MARROCOS!
Output: Olá o seu input é um palindromo
``` 

### Challenge #2 - Sequência de Fibonacci

Fibonacci é uma sequência de números inteiros começando normalmente por 0 e 1, na qual, cada termo subsequente corresponde à soma dos dois números anteriores. Os números de Fibonacci são, portanto, os números que compõem a seguinte sequência: `0, 1, 1, 2, 3, 5, 8, 13, 21, ...`.

Em matemática, está sequência é definida [**recursivamente**](https://pt.wikipedia.org/wiki/Recurs%C3%A3o) pela fórmula abaixo, sendo o primeiro termo ![first term](https://latex.codecogs.com/gif.latex?F_%7B1%7D%3D1):

- ![formula](https://latex.codecogs.com/gif.latex?F_%7Bn%7D%3DF_%7Bn-1%7D&plus;F_%7Bn-2%7D)

valores iniciais:
- ![initial value](https://latex.codecogs.com/gif.latex?F_%7B1%7D%3D%201%2C%20F_%7B2%7D%20%3D%202)

O seu programa deve receber o número *n* e gerar uma sequencia de fibonacci.

```sh
input: 10
output: Fib: 0 1 1 2 3 5 8 13 21 34
```

### Challenge #3 - Probabilidade

Dada uma moeda justa que é lançada `N` vezes, a tarefa é determinar a probabilidade de forma que não ocorram duas *insígnias* consecutivamente.

```sh
input: N = 2
output: 0.75
```
> T: Numeros, H: Insígnia 👇🏾

Quando a moeda é lançada 2 vezes, os resultados possíveis são `TH, HT, TT, HH`. Como em 3 de 4 resultados, as *insígnias* não ocorrem juntas. Portanto, a probabilidade requerida é `(3/4)` ou `0.75`

**Dica importante:** use o conceito aprendido quando no exercício sobre Fibonacci para resolver este exercício.

<hr/>

> *Let's hack it. Boa sorte ~ Fimba code: 22, May, 2020*
