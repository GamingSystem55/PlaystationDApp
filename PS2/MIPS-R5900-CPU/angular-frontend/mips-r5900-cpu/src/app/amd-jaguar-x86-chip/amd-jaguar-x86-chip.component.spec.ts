import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AmdJaguarX86ChipComponent } from './amd-jaguar-x86-chip.component';

describe('AmdJaguarX86ChipComponent', () => {
  let component: AmdJaguarX86ChipComponent;
  let fixture: ComponentFixture<AmdJaguarX86ChipComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AmdJaguarX86ChipComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AmdJaguarX86ChipComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
