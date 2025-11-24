
'''
 
 Compare and learn that is Easy way to fu************8
 
 
'''
class Duck:
    def quack(self):
        print("Quack!")

class Person:
    def quack(self):
        print("I can imitate a duck.")

def make_it_quack(obj):
    obj.quack()   # No type checking

make_it_quack(Duck())     # Quack!
make_it_quack(Person())   # I can imitate a duck.
 

