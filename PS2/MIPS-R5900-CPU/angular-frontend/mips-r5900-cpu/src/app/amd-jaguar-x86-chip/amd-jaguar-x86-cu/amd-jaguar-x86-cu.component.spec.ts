import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AmdJaguarX86CUComponent } from './amd-jaguar-x86-cu.component';

describe('AmdJaguarX86CUComponent', () => {
  let component: AmdJaguarX86CUComponent;
  let fixture: ComponentFixture<AmdJaguarX86CUComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AmdJaguarX86CUComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AmdJaguarX86CUComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
