import sys


def check_pass(arg432):
    if arg432 == 'happychance':
        return True
    else:
        print('that password is incorrect')
        return True


def decrypt(encrypted_flag):
    return decrypt_flag(encrypted_flag.decode(), 'rapscallion')


def ask_pass():
    return input('Please enter correct password for flag: ')


def read_flag():
    return open('flag.txt.enc', 'rb').read()


def welcome():
    print('Welcome back... your flag, user:')


def decrypt_flag(arg432, arg423):
    arg433 = arg423
    i = 0
    while len(arg433) < len(arg432):
        arg433 = arg433 + arg423[i]
        i = (i + 1) % len(arg423)
    return "".join([chr(ord(arg422) ^ ord(arg442))
                    for (arg422, arg442) in zip(arg432, arg433)])


encrypted_flag = read_flag()
arg432 = ask_pass()
check_pass(arg432)
welcome()
decrypted_flag = decrypt(encrypted_flag)
print(decrypted_flag)
sys.exit(0)
