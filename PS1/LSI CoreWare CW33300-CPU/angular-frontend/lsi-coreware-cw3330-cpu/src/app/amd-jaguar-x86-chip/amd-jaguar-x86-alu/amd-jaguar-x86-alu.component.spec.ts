import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AmdJaguarX86ALUComponent } from './amd-jaguar-x86-alu.component';

describe('AmdJaguarX86ALUComponent', () => {
  let component: AmdJaguarX86ALUComponent;
  let fixture: ComponentFixture<AmdJaguarX86ALUComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AmdJaguarX86ALUComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AmdJaguarX86ALUComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
