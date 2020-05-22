<h1 align="center" style="color:">Desafio de Algoritmos básicos: #1 </h1>
  
## Introdução 📜

Este desafio foi desenhado para programadores iniciantes, e também os mais experientes que desejam acrescentar os seus conhecimentos acerca da resolução de alguns algoritmos em programação.  As soluções não estarão restritas à nenhuma linguagem de programação específica, os nossos **moderadores** encorajam todos os participantes a usar as linguagens de programação de sua preferência. 

## Submissão 🚀

As soluções desenvolvidas devem ser partilhadas no canal [**#code-drop** no discord](https://discord.gg/XDPbSUN), posteriormente, as melhores soluções serão colocadas neste repo e os devs receberão kudos no twitter ou outra rede social de sua preferência. 

**Dicas importantes**

1. Não te esqueças de formatar o teu código antes de o enviar à communidade, para mais informações visite => [discord](https://support.discord.com/hc/en-us/articles/210298617-Markdown-Text-101-Chat-Formatting-Bold-Italic-Underline).
2. Use funções no teu programa sempre que necessário para a reutilização de códigos e simplificação dos programas. 
3. Resolva os teus exercícios streaming live (ajuda-te a ganhar mais pontos)
4. Desenvolva e partilhe o seu pseudocódigo(algoritimo) junto da tua submissão.


## Challenge 🥋

Os challenges (desafios) começarão do mais simples ao mais arduo, mas poderás resolver-los de forma aleatória. Esse challenge será composto por 4 problemas sendo que os mesmos ajudarão o participante à desenvolver algum skill específico.

### Challenge #1 - Soma de números inteiros 🧊

Dado um número inteiro, retorne a soma de todos os dígitos.

```sh
input: 124 
output: 7
```

### Challenge #2 - Palíndromo 💤

Palíndromos é um número/palavra (ou um conjunto de números/palavras) que **lendo da esquerda para a direita ou da direita para esquerda são iguais**. Para este challenge, o teu programa deve receber um input do usuário e definir se este input é ou não um palindromo. 

```sh
Input: SOCORRAM-ME! SUBI NO ÔNIBUS EM MARROCOS!
Output: Olá o seu input é um palindromo
``` 

### Challenge #3 - Sequência de Fibonacci 🔢

Fibonacci é uma sequência de números inteiros começando normalmente por 0 e 1, na qual, cada termo subsequente corresponde à soma dos dois números anteriores. Os números de Fibonacci são, portanto, os números que compõem a seguinte sequência: `0, 1, 1, 2, 3, 5, 8, 13, 21, ...`.

Em matemática, está sequência é definida [**recursivamente**](https://pt.wikipedia.org/wiki/Recurs%C3%A3o) pela fórmula abaixo, sendo o primeiro termo ![first term](https://latex.codecogs.com/gif.latex?F_%7B1%7D%3D1):

- ![formula](https://latex.codecogs.com/gif.latex?F_%7Bn%7D%3DF_%7Bn-1%7D&plus;F_%7Bn-2%7D)

valores iniciais:
- ![initial value](https://latex.codecogs.com/gif.latex?F_%7B1%7D%3D%201%2C%20F_%7B2%7D%20%3D%202)

O seu programa deverá receber o número *n* e gerar uma sequência de fibonacci.

```sh
input: 10
output: Fib: 0 1 1 2 3 5 8 13 21 34
```

### Challenge #4 - Probabilidade 🎲

Dada uma moeda justa que é lançada `N` vezes, a tarefa é determinar a probabilidade de forma que não ocorram duas *insígnias* consecutivamente.

```sh
input: N = 2
output: 0.75
```
> T: Numeros, H: Insígnia 👇🏾

Quando a moeda é lançada 2 vezes, os resultados possíveis são `TH, HT, TT, HH`. Como em 3 de 4 resultados, as *insígnias* não ocorrem juntas. Portanto, a probabilidade requerida é `(3/4)` ou `0.75`

**Dica importante:** use o conceito aprendido aquando da resolução do exercício sobre Fibonacci para resolver este exercício.

<hr/>

> *Let's hack it. Boa sorte ~ Fimba code: 22, May, 2020*
