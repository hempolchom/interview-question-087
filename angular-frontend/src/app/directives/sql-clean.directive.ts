import { Directive, HostListener } from '@angular/core';

@Directive({
  selector: '[appSqlClean]',
  standalone: true
})
export class SqlCleanDirective {
  private regex = new RegExp(/[^a-zA-Z0-9_-]/g);

  @HostListener('input', ['$event'])
  onInput(event: InputEvent): void {
    const input = event.target as HTMLInputElement;
    if (input && input.value) {
      input.value = input.value.replace(this.regex, '');
      input.dispatchEvent(new Event('input'));
    }
  }
}
