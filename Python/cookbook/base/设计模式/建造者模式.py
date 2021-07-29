from abc import ABCMeta, abstractmethod


class Director(metaclass=ABCMeta):
    def __init__(self):
        self._builder = None
    def set_builder(self, builder):
        self._builder = builder
    @abstractmethod
    def construct(self, field_list):
        pass
    def get_constructed_object(self):
        return self._builder.constructed_object

class AbstractFormBuilder(metaclass=ABCMeta):
    def __init__(self):
        self.constructed_object = None
    @abstractmethod
    def add_text_field(self, field_fict):
        pass
    @abstractmethod
    def add_checkbox(self, checkbox_dict):
        pass
    @abstractmethod
    def add_button(self, button_dict):
        pass

class HtmlForm:
    def __init__(self):
        self.field_list = []
    def __repr__(self):
        return '<form>{}</form>'.format("".join(self.field_list))

class HtmlFormBuilder(AbstractFormBuilder):
    def __init__(self):
        self.constructed_object = HtmlForm()
    def add_text_field(self, field_dict):
        self.constructed_object.field_list.append(
            '{}:<br><input type="text" name="{}"></br>'.format(field_dict['label'], field_dict['field_name'])
        )
    def add_checkbox(self, checkbox_dict):
        self.constructed_object.field_list.append(
            '<label>{label}</label><input type="checkbox" id="{id}" value="{value}">'.format(
                id=checkbox_dict['field_id'], 
                value=checkbox_dict['value'],
                label=checkbox_dict['label']
            )           
        )
    def add_button(self, button_dict):
        self.constructed_object.field_list.append(
            '<input type="button" value="{}" />'.format(button_dict['text'])
        )

class FormDirector(Director):
    def __init__(self):
        super().__init__()
    def construct(self, field_list):
        for field in field_list:
            fieldtype = field['field_type']
            if fieldtype == 'text_field':
                self._builder.add_text_field(field)
            elif fieldtype == 'checkbox':
                self._builder.add_checkbox(field)
            elif fieldtype == 'button':
                self._builder.add_button(field)

if __name__ == "__main__":
    director = FormDirector()
    html_form_builder = HtmlFormBuilder()
    director.set_builder(html_form_builder)

    field_list = [
        {
            'field_type': 'text_field',
            'label': 'some text here',
            'field_name': 'field one'
        },
        {
            'field_type': 'checkbox',
            'field_id': 'check_it',
            'value': '1',
            'label': 'check for on'
        },
        {
            'field_type': 'button',
            'text': 'button here'
        }
    ]
    director.construct(field_list)
    print('<html><body>{0!r}</body></html>'.format(director.get_constructed_object()))