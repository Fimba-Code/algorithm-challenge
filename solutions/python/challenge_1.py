# 1. Receber input do usuário
# input é recebido como string
num = input("input: ")

# 2. Verificar se o input é um número inteiro
try:
    # Tentar converter o input para inteiro
    int(num)
except ValueError:
    # Em caso de erro sair do programa
    print("Introduza um número inteiro")
    exit(0)


#3. Remover o sinal '-' se o número for negativo

if "-" in num:
    num = num.replace("-","")


# 4. Percorrer cada 'digito' do número e acumular a soma 
# dentro da var 'result'
result = 0

for digit in num:
    # cada digito é primeiro convertido para int e 
    #só depois é acrescentado à soma
    result += int(digit)
 
# 5. Mostrar resultado ao usuário    
print(f"output: {result}")
