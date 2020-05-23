# Challenge #4 - Probabilidade
#Constantes que representam as diferentes faces da moeda
T = "T"
H = "H"

# função para calcular todas permutações possíveis
# Todas as permutações são postas num array
#limit: é o número de lançamentos e limita o 
# tamanho da string. Ex: p/ limit = 2, as strings geradas são
# HT,TT,TH,HH
# perms: é o array que armazena as permutações
# current: armazena o estado da string que formará a 
#próxima permutação. inicialmente é uma string vazia
#depth: controla a profundidade da recursão
def get_perms(limit,perms,current="",depth=0):
    # 1. Verificar se atingimos o caso base da recursão
    if depth == limit:
        #se sim, a string actual é uma permutação válida
        # e é adicionada ao array de permutações 
        perms.append(current)
        return

    # 2. Chamar de forma recursiva get_perms modificando 
    # a string actual com as variações possíveis (T e H).
    get_perms(limit,perms,current+T,depth+1)   
    get_perms(limit,perms,current+H,depth+1)


# 1. receber o número de lançamentos do usuário
flips = int(input("input: "))


# 2. Gerar todas as permutações de 'H' e 'T' com base no número de
# lançamentos. Ex: para flips = 3 teriamos: TTT, THT, TTH ...
#o número total de resultados é sempre  2^flips
results = []
#a função get_perms gera os resultados e adiciona-os ao array
get_perms(flips,results)


# 3. Contar o número de resultados com dois 'H's 
#seguidos
count = 0

for result in results:
    #verifica cada resultado
    # e procura pelo padrão 'HH' em qualquer
    # posição na string 
    if "HH" in result:
        count +=1

# 4. Calcular e mostrar a probabilidade para o usuário
# Lembrando que a probabilidade = x/y
# onde x -> acontecimentos favoráveis (quantidade de resultados sem 'HH') = tamanho(resultados) - quantidade_de_res_com_HH
# onde y -> todos acontecimentos possíveis (quantidade de resultados) = tamanho(resultados)  
print(f"ouput: {(len(results)-count)/len(results)}")
