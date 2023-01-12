from jproperties import Properties

class Grammar:

    def __init__(self, filename_):
        
        self._file = open(file=filename_, mode='rt')

        self._properties = Properties()
        self._properties.properties = {"nonterminals": [], "alphabet": [], "productions": [], "initial_state": []}
        self._properties.load(self._file, encoding=None)

        self._properties.properties["nonterminals"] = self._properties.properties["nonterminals"].strip().split(" ")
        self._properties.properties["alphabet"] = self._properties.properties["alphabet"].strip().split(" ")
        self._properties.properties["initial_state"] = self._properties.properties["initial_state"]
    
        raw_productions_list = self._properties.properties["productions"].strip().split(" ~ ")

        self._properties.properties["productions"] = []

        for pr in raw_productions_list:
            
            production = pr.strip().split("->", 1)

            new_pr = Production(production[0].strip().split(" "), production[1])

            self._properties.properties["productions"].append(new_pr)

        print(self.check_CFG())


    @property
    def nonterminals(self):
        return self._properties.properties["nonterminals"]
    

    @property
    def alphabet(self):
        return self._properties.properties["alphabet"]

    @property
    def initial_state(self):
        return self._properties.properties["initial_state"]
            
    @property
    def productions(self):
        return self._properties.properties["productions"]
    @property
    def str_productions(self):
        local_list=""
        for elem in self._properties.properties["productions"]:
            local_list += str(elem) + "\n"
        return local_list
            

    def productions_for_id(self,id_):
        local_list=[]
        for pr in self._properties.properties["productions"]:
            if(pr.id[0]==id_):
                local_list.append(pr)
        return local_list
    

    def check_CFG(self):

        check_initial=False

        for pr in self._properties.properties["productions"]:

            if self._properties.properties["initial_state"] in pr.id:
                check_initial=True
            if len(pr.id)>1 or pr.id[0] not in self._properties.properties["nonterminals"]:
                return False

            for elem in pr.list:
                for elem2 in elem:
                    if elem2 not in self._properties.properties["alphabet"] and elem2 not in self._properties.properties["nonterminals"]:
                        print(elem2)
                        return False

        if not check_initial:
            return False
        return True
        
    def __str__(self):
        return("non-terminals: "+str(self.nonterminals)+"\n"+
            "alphabet: "+str(self.alphabet)+"\n"+
            "initial_state: "+str(self.initial_state)+"\n"
            "productions:\n"+str(self.str_productions)+"\n"
            )           


class Production:

    def __init__(self, id_: list, list_: str):
        self._id = id_
        
        self._list  = list_.strip().split("|")

        for index in range(len(self._list)):

            self._list[index] = self._list[index].strip().split(" ")

    def __str__(self):
        return f"{self._id} -> {self._list}"
    
    
    @property
    def id(self):
        return self._id
    
    @property
    def list(self):
        return self._list
    