## https://www.techbeamers.com/class-definitions-in-python

## 1. Simple Class Definition
class Pen:
    def __init__(self, make, color):
        self.make = make
        self.color = color

    def write(self):
        print(f"The {self.make} pen has {self.color} ink.")

my_pen = Pen("Pilot", "Blue")
print(my_pen.make)
my_pen.write()

## 2. Class Methods
class Pen:
    @classmethod
    def create(cls, make, color):
        return cls(make, color)

author_pen = Pen.create("Hauser", "Green")

## 3. Class Variables
class Pen:
    num_pens = 0

    def __init__(self, make, color):
        self.make = make
        self.color = color
        Pen.num_pens += 1

print(Pen.num_pens)

## 4. Inheritance
class Ballpoint(Pen):
    pass

## 5. Polymorphism
class Pen:
    def __init__(self, make, color):
        self.make = make
        self.color = color

    def write(self):
        print(f"The {self.make} pen has {self.color} ink.")

    @classmethod
    def create(cls, make, color):
        return cls(make, color)

class Ballpoint(Pen):
    def write(self):
        print(f"The {self.make} Ballpoint pen has {self.color} ink.")

pen = Pen.create("Hauser", "Red")
point = Ballpoint.create("Linc", "Orange")

# Call the write() method on both the pen and the point
pen.write()
point.write()

## 6. Unique Example Code Snippets
class Fan:
    fan_count = 0

    def __init__(self, brand, type):
        self.brand = brand
        self.type = type
        self.is_on = False
        Fan.fan_count += 1

    def turn_on(self):
        if not self.is_on:
            self.is_on = True
            print(f"The {self.brand} {self.type} fan is now turned on.")
        else:
            print(f"The {self.brand} {self.type} fan is already on.")

    def turn_of(self):
        if self.is_on:
            self.is_on = False
            print(f"The {self.brand} {self.type} fan is now turned off.")
        else:
            print(f"The {self.brand} {self.type} fan is already off.")

    def display_info(self):
        print(f"Brand: {self.brand}, Type: {self.type}, Is On: {self.is_on}")


# create
my_fan = Fan("Dyson", "Tower")

# turn on the fan
my_fan.turn_on()

# try to turn on the same fan again
my_fan.turn_on()

# turn off the fan
my_fan.turn_off()

# display fan information
my_fan.display_info()

print(f"Total number of fands: {Fan.fan_count}")